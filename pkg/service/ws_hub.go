package service

import chat "github.com/MerBasNik/rndmCoffee"

type Hub chat.Hub

func NewHub() *Hub {
	return &Hub{
		Chats:      make(map[string]*chat.ChatList),
		Register:   make(chan *chat.Client),
		Unregister: make(chan *chat.Client),
		Broadcast:  make(chan *chat.ChatItem),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Chats[cl.RoomId]; ok {
				r := h.Chats[cl.RoomId]

				if _, ok := r.UsersId[cl.Id]; !ok {
					r.UsersId[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Chats[cl.RoomId]; ok {
				if _, ok := h.Chats[cl.RoomId].UsersId[cl.Id]; ok {
					if len(h.Chats[cl.RoomId].UsersId) != 0 {

						h.Broadcast <- &chat.ChatItem{
							Description: "user left the chat",
							Chatlist_id: cl.RoomId,
							Username:    cl.Username,
						}
					}
					delete(h.Chats[cl.RoomId].UsersId, cl.Id)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Chats[m.Chatlist_id]; ok {
				for _, cl := range h.Chats[m.Chatlist_id].UsersId {
					cl.Message <- m
				}
			}
		}
	}
}
