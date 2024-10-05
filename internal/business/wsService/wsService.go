package wsservice

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"
)

type WSService struct {
	wsDao IDataAccessWS
}

type IDataAccessWS interface {
	GetUser(ctx context.Context, GUID string) (*models.User, error)
	GetChatroom(ctx context.Context, chatroomID string) (*models.Chatroom, error)
}

func New(wsdao IDataAccessWS) *WSService {
	return &WSService{
		wsDao: wsdao,
	}
}

func (w *WSService) GetUser(ctx context.Context, GUID string) (*models.User, error) {
	user, err := w.wsDao.GetUser(ctx, GUID)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return user, nil
}

func (w *WSService) GetChatroom(ctx context.Context, chatroomID string) (*models.Chatroom, error) {
	chatroom, err := w.wsDao.GetChatroom(ctx, chatroomID)
	slog.Debug(fmt.Sprintf("%v", chatroom))
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}


	return chatroom, nil
}
