package socket

import (
	"context"
	"log"
	"realtime-chatapp/internal/domain"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan domain.MessageResponse
	UserID int
}

func (c *Client) ReadPump(u domain.MessageUsecase) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var req domain.MessageRequest
		err := c.Conn.ReadJSON(&req)
		if err != nil {
			log.Printf("error read json: %v", err)
			break
		}
		res, err := u.SendMessage(context.Background(), req, c.UserID)
		if err != nil {
			log.Printf("gagal simpan pesan: %v", err)
			continue
		}
		c.Hub.Broadcast <- res
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := c.Conn.WriteJSON(message)
		if err != nil {
			log.Println("Error write JSON:", err)
			return
		}
	}
}
