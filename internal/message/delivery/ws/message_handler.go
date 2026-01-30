package ws

import (
	"net/http"
	"realtime-chatapp/internal/domain"
	"realtime-chatapp/internal/helper"
	"realtime-chatapp/internal/message/delivery/ws/socket"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageHandlerImpl struct {
	Usecase domain.MessageUsecase
	hub     *socket.Hub
}

func NewMessageHandler(u domain.MessageUsecase, h *socket.Hub) domain.MessageHandler {
	return &MessageHandlerImpl{
		Usecase: u,
		hub:     h,
	}
}

func (m *MessageHandlerImpl) GetChatHistory(c *gin.Context) {
	senderID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	receiverID, _ := strconv.Atoi(c.Query("receiver_id"))

	if receiverID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receiver_id is required"})
		return
	}
	history, err := m.Usecase.GetChatHistory(c.Request.Context(), senderID.(int), receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

func (m *MessageHandlerImpl) Connect(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token diperlukan"})
		return
	}
	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &socket.Client{
		Hub:    m.hub,
		Conn:   conn,
		Send:   make(chan domain.MessageResponse, 256),
		UserID: claims.UserID,
	}
	m.hub.Register <- client
	go client.WritePump()
	go client.ReadPump(m.Usecase)
}
