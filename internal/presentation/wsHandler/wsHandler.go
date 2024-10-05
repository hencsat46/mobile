package wsHandler

import (
	"context"
	"log"
	"log/slog"
	"time"

	"hackathon/internal/presentation/entities"
	hubmanager "hackathon/internal/presentation/hubManager"
	messagehttphandler "hackathon/internal/presentation/messageHTTPhandler"
	"hackathon/models"

	"github.com/gofiber/contrib/websocket"
)

type WSHandler struct {
	hubManager  *hubmanager.HubManager
	WsBusiness  IBusinessWS
	msgBusiness messagehttphandler.IBusinessMessage
}

type IBusinessWS interface {
	GetUser(ctx context.Context, GUID string) (*models.User, error)
	GetChatroom(ctx context.Context, chatroomID string) (*models.Chatroom, error)
}

func New(wsBusiness IBusinessWS, hubmngr *hubmanager.HubManager, msgBusiness messagehttphandler.IBusinessMessage) *WSHandler {
	return &WSHandler{
		hubManager:  hubmngr,
		WsBusiness:  wsBusiness,
		msgBusiness: msgBusiness,
	}
}

func (h *WSHandler) HandleWS(c *websocket.Conn) {
	guid := c.Params("GUID")
	cid := c.Params("cid")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	user, err := h.WsBusiness.GetUser(ctx, guid)
	if err != nil || user == nil {
		slog.Debug(err.Error() + " or no such user exists")
		return
	}

	chatroom, err := h.WsBusiness.GetChatroom(ctx, cid)
	if err != nil || chatroom == nil {
		slog.Debug(err.Error() + " or no such chatroom exists")
		return
	}

	log.Println(chatroom)
	h.hubManager.AddParticipant(c, cid, guid)
	h.listenUserMessage(c, cid, guid)
}

func (h *WSHandler) listenUserMessage(c *websocket.Conn, cid, guid string) {
	for {
		msg := &entities.Message{}
		if err := c.ReadJSON(msg); err != nil {
			slog.Debug(err.Error())
			h.hubManager.DeleteParticipant(c, cid, guid)
			return
		}

		message := models.Message{
			MessageId:  msg.MessageID,
			ChatroomId: msg.ChatroomID,
			SenderGUID: msg.GUID,
			Content:    msg.Content,
			Image:      msg.Image,
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()

		messageID, name, err := h.msgBusiness.CreateMessage(ctx, message)
		if err != nil {
			slog.Debug(err.Error())
		}

		msg.MessageID = messageID
		msg.SenderName = name

		h.hubManager.SendMessage(msg)
	}
}
