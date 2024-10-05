package hubloaderservice

import (
	"context"
	"hackathon/models"
	"log"
	"log/slog"
)

type HubLoader struct {
	hubdao IDataAccessLoader
}

type IDataAccessLoader interface {
	GetChatrooms(context.Context) ([]models.Chatroom, error)
}

func New(hubdao IDataAccessLoader) *HubLoader {
	return &HubLoader{
		hubdao: hubdao,
	}
}

func (h *HubLoader) GetChatrooms(ctx context.Context) ([]models.Chatroom, error) {
	rooms, err := h.hubdao.GetChatrooms(ctx)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	log.Println(rooms)

	return rooms, nil
}
