package chatroomservice

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"

	"github.com/google/uuid"
)

type ChatroomService struct {
	ChatroomDao IDataAccessChatroom
}

type IDataAccessChatroom interface {
	GetChatrooms(ctx context.Context) ([]models.Chatroom, error)
	CreateChatroom(ctx context.Context, chatroom models.Chatroom) error
	UpdateChatroom(ctx context.Context, chatroomID, chatroomName string) error
	DeleteChatroom(ctx context.Context, ownerGUID, chatroomID string) error
}

func New(chatroomdao IDataAccessChatroom) *ChatroomService {
	return &ChatroomService{
		ChatroomDao: chatroomdao,
	}
}

func (b *ChatroomService) CreateChatroom(ctx context.Context, chatroom models.Chatroom) (string, error) {
	slog.Debug(fmt.Sprintf("creating chatroom: %v", chatroom))

	chatroom.ChatroomId = uuid.NewString()

	if err := b.ChatroomDao.CreateChatroom(ctx, chatroom); err != nil {
		slog.Debug(err.Error())
		return "", err
	}
	return chatroom.ChatroomId, nil
}

func (b *ChatroomService) UpdateChatroom(ctx context.Context, chatroomID, chatroomName string) error {
	slog.Debug(fmt.Sprintf("upodating chatroom: %v", chatroomID))
	if err := b.ChatroomDao.UpdateChatroom(ctx, chatroomID, chatroomName); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *ChatroomService) DeleteChatroom(ctx context.Context, ownerGUID, chatroomID string) error {
	slog.Debug(fmt.Sprintf("deleting chatroom: %v", chatroomID))
	if err := b.ChatroomDao.DeleteChatroom(ctx, ownerGUID, chatroomID); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (b *ChatroomService) GetChatrooms(ctx context.Context) ([]models.Chatroom, error) {
	slog.Debug("getting chatrooms")
	chatrooms, err := b.ChatroomDao.GetChatrooms(ctx)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return chatrooms, nil
}
