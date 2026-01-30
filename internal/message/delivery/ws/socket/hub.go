package socket

import (
	"log"
	"realtime-chatapp/internal/domain"
)

type Hub struct {
	Clients    map[int]*Client
	Broadcast  chan domain.MessageResponse
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int]*Client),
		Broadcast:  make(chan domain.MessageResponse),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserID] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			log.Printf("Hub mendistribusikan pesan dari %d ke %d", message.Sender.ID, message.Receiver.ID)
			if receiver, ok := h.Clients[message.Receiver.ID]; ok {
				select {
				case receiver.Send <- message:
				default:
					close(receiver.Send)
					delete(h.Clients, message.Receiver.ID)
				}
			}
			if sender, ok := h.Clients[message.Sender.ID]; ok {
				select {
				case sender.Send <- message:
				default:
					close(sender.Send)
					delete(h.Clients, message.Sender.ID)
				}
			}
		}
	}
}
