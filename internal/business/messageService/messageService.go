package messageservice

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"

	"github.com/beevik/guid"
)

type MessageService struct {
	MessageDataAccess IDataAccessMessage
}

type IDataAccessMessage interface {
	FetchMessagesForChatroom(ctx context.Context, chatroomID string) ([]models.Message, error)
	CreateMessage(ctx context.Context, message models.Message) (string, error)
	UpdateMessage(ctx context.Context, newContent, messageID, chatroomID string) error
	DeleteMessage(ctx context.Context, message models.Message) error
}

func New(msgDao IDataAccessMessage) *MessageService {
	return &MessageService{
		MessageDataAccess: msgDao,
	}
}

func (b *MessageService) FetchMessagesForChatroom(ctx context.Context, chatroomID string) ([]models.Message, error) {
	messages, err := b.MessageDataAccess.FetchMessagesForChatroom(ctx, chatroomID)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return messages, nil
}

func (b *MessageService) CreateMessage(ctx context.Context, message models.Message) (string, string, error) {
	message.MessageId = guid.NewString()

	name, err := b.MessageDataAccess.CreateMessage(ctx, message)
	if err != nil {
		slog.Debug(err.Error())
		return "", "", err
	}
	return message.MessageId, name, nil
}

func (b *MessageService) UpdateMessage(ctx context.Context, newContent, messageID, chatroomID string) error {
	slog.Debug(fmt.Sprintf("updating message: %v, %v", messageID, chatroomID))
	if err := b.MessageDataAccess.UpdateMessage(ctx, newContent, messageID, chatroomID); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *MessageService) DeleteMessage(ctx context.Context, message models.Message) error {
	if err := b.MessageDataAccess.DeleteMessage(ctx, message); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
