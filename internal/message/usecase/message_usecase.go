package usecase

import (
	"context"
	"realtime-chatapp/internal/domain"
	"time"
)

type MessageUsecaseImpl struct {
	repo domain.MessageRepository
}

func NewMessageUsecase(repo domain.MessageRepository) domain.MessageUsecase {
	return &MessageUsecaseImpl{repo: repo}
}
func (m *MessageUsecaseImpl) SendMessage(ctx context.Context, req domain.MessageRequest, senderID int) (domain.MessageResponse, error) {
	msg := domain.Message{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
		CreatedAt:  time.Now(),
	}

	if err := m.repo.Save(ctx, msg); err != nil {
		return domain.MessageResponse{}, err
	}
	history, err := m.repo.GetChatHistory(ctx, senderID, req.ReceiverID)
	if err != nil || len(history) == 0 {
		return domain.ToMessageResponse(msg, "", ""), nil
	}
	return history[len(history)-1], nil

}
func (m *MessageUsecaseImpl) GetChatHistory(ctx context.Context, senderID, receiverID int) ([]domain.MessageResponse, error) {
	history, err := m.repo.GetChatHistory(ctx, senderID, receiverID)
	if err != nil {
		return []domain.MessageResponse{}, nil
	}
	if history == nil {
		return []domain.MessageResponse{}, nil
	}
	return history, nil
}
