package hubmanager

import (
	"context"
	"log"
	"log/slog"
	"sync"
	"time"

	"hackathon/internal/presentation/entities"
	"hackathon/models"

	"github.com/gofiber/contrib/websocket"
)

type HubManager struct {
	sync.RWMutex
	hub    entities.WSHub
	loader ILoader
}

type ILoader interface {
	GetChatrooms(context.Context) ([]models.Chatroom, error)
}

func New(loader ILoader) *HubManager {
	return &HubManager{
		loader: loader,
		hub:    make(entities.WSHub),
	}
}

func (h *HubManager) LoadChatroomToHub(room *entities.WSRoom) {
	h.Lock()
	defer h.Unlock()
	room.Participants = make(map[string]*websocket.Conn)
	h.hub[room.CID] = room
}

func (h *HubManager) LoadAllChatroomsToHub() {
	h.Lock()
	defer h.Unlock()
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	rooms, err := h.loader.GetChatrooms(ctx)
	if err != nil {
		slog.Debug(err.Error())
	}
	rms := make([]entities.WSRoom, 0, len(rooms))
	
	for _, r := range rooms {
		rms = append(rms, entities.WSRoom{CID: r.ChatroomId, Participants: make(map[string]*websocket.Conn)})
	}
	log.Println("rms", rms)
	for _, r := range rms {
		h.hub[r.CID] = &r
	}
	log.Println(h.hub["ca7bbc89-b569-44c1-b00e-c3528acfad1a"])
}

func (h *HubManager) AddParticipant(c *websocket.Conn, cid, guid string) {
	h.Lock()
	defer h.Unlock()
	log.Println("add participant: ", h.hub["ca7bbc89-b569-44c1-b00e-c3528acfad1a"])
	h.hub[cid].Participants[guid] = c
	log.Println("after adding")
}

func (h *HubManager) DeleteParticipant(c *websocket.Conn, cid, guid string) {
	h.Lock()
	defer h.Unlock()
	delete(h.hub[cid].Participants, guid)
}

func (h *HubManager) SendMessage(msg *entities.Message) {
	h.RLock()
	defer h.RUnlock()
	chatroomID := msg.ChatroomID
	room := h.hub[chatroomID]

	for GUID, conn := range room.Participants {
		if GUID != msg.GUID {
			if err := conn.WriteJSON(msg); err != nil {
				slog.Error(err.Error())
			}
		}
	}
}
