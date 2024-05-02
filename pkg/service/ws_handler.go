package service

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HandlerWS struct {
	hub *Hub
}

func NewHandlerWS(h *Hub) *HandlerWS {
	return &HandlerWS{
		hub: h,
	}
}

type CreateRoomReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *HandlerWS) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Chats[req.Id] = &chat.ChatList{
		Id:      req.Id,
		Title:   req.Name,
		UsersId: make(map[string]*chat.Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		orgin := r.Header.Get("Origin")
		return orgin == "http://localhost:3000"
	},
}

func (h *HandlerWS) JoinRoom(c *gin.Context, CreateItem func(int, int, string, string, string) (int, error)) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomId := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		&chat.Client{
			Conn:     conn,
			Message:  make(chan *chat.ChatItem, 10),
			Id:       clientId,
			RoomId:   roomId,
			Username: username,
		},
	}

	var description = "Новый пользователь зашел в комнату"
	clientIdNum, err := strconv.Atoi(clientId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomIdNum, err := strconv.Atoi(roomId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idChatItem, err := CreateItem(clientIdNum, roomIdNum, username, description, roomId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := &chat.ChatItem{
		Description: description,
		Id:          idChatItem,
		Username:    username,
		Chatlist_id: roomId,
	}

	h.hub.Register <- cl.Client
	h.hub.Broadcast <- m

	go cl.writeChatItem(c, CreateItem)
	cl.ReadChatItem(h.hub, c, CreateItem)
}
