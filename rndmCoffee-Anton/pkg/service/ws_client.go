package service

import (
	"net/http"
	"strconv"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	*chat.Client
}

func (c *Client) writeChatItem(context *gin.Context, CreateItem func(int, int, string, string, string) (int, error)) {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		userId, err := strconv.Atoi(c.Id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		listId, err := strconv.Atoi(message.Chatlist_id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		CreateItem(userId, listId, message.Username, message.Description, message.Chatlist_id)

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadChatItem(hub *Hub, context *gin.Context, CreateItem func(int, int, string, string, string) (int, error)) {
	defer func() {
		hub.Unregister <- c.Client
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			break
		}

		userId, err := strconv.Atoi(c.Id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		listId, err := strconv.Atoi(c.RoomId)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		CreateItem(userId, listId, c.Username, string(m), c.RoomId)

		msg := &chat.ChatItem{
			Description: string(m),
			Chatlist_id: c.RoomId,
			Username:    c.Username,
		}

		hub.Broadcast <- msg
	}
}
