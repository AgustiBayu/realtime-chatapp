package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Content    string
	CreatedAt  time.Time
}

type UserShortResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MessageRequest struct {
	ReceiverID int    `json:"receiver_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type MessageResponse struct {
	ID        int               `json:"id"`
	Sender    UserShortResponse `json:"sender"`
	Receiver  UserShortResponse `json:"receiver"`
	Content   string            `json:"content"`
	CreatedAt string            `json:"created_at"`
}

type MessageRepository interface {
	Save(ctx context.Context, message Message) error
	GetChatHistory(ctx context.Context, senderID, receiverID int) ([]MessageResponse, error)
}

type MessageUsecase interface {
	SendMessage(ctx context.Context, req MessageRequest, senderID int) (MessageResponse, error)
	GetChatHistory(ctx context.Context, senderID, receiverID int) ([]MessageResponse, error)
}

type MessageHandler interface {
	GetChatHistory(c *gin.Context)
	Connect(c *gin.Context)
}

func ToMessageResponse(msg Message, senderName, receiverName string) MessageResponse {
	return MessageResponse{
		ID: msg.ID,
		Sender: UserShortResponse{
			ID:   msg.SenderID,
			Name: senderName,
		},
		Receiver: UserShortResponse{
			ID:   msg.ReceiverID,
			Name: receiverName,
		},
		Content:   msg.Content,
		CreatedAt: msg.CreatedAt.Format(time.RFC3339),
	}
}
