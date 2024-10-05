package messagehttphandler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	e "hackathon/exceptions"
	"hackathon/internal/presentation/entities"
	"hackathon/models"

	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	MessageBusiness IBusinessMessage
}

type IBusinessMessage interface {
	FetchMessagesForChatroom(ctx context.Context, chatroomID string) ([]models.Message, error)
	CreateMessage(ctx context.Context, message models.Message) (string, string, error)
	UpdateMessage(ctx context.Context, newContent, messageID, chatroomID string) error
	DeleteMessage(ctx context.Context, message models.Message) error
}

func New(messageBusiness IBusinessMessage) *MessageHandler {
	return &MessageHandler{
		MessageBusiness: messageBusiness,
	}
}

func (h *MessageHandler) FetchMessagesForChatroom(c *fiber.Ctx) error {
	chatroomId := c.Params("cid")
	slog.Debug(fmt.Sprintf("fetch chatroom messages endpoint called: %v", chatroomId))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	messages, err := h.MessageBusiness.FetchMessagesForChatroom(ctx, chatroomId)
	if err != nil {
		slog.Debug(err.Error())
		if errors.Is(err, e.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(entities.Response{
				Error:   e.ErrNotFound.Error(),
				Content: nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: messages,
	})
}

func (h *MessageHandler) UpdateMessage(c *fiber.Ctx) error {
	var request entities.MessageDTO

	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update message endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.MessageBusiness.UpdateMessage(ctx, request.Content, request.MessageId, request.ChatroomID); err != nil {
		slog.Debug(err.Error())
		if errors.Is(err, e.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(entities.Response{
				Error:   e.ErrNotFound.Error(),
				Content: nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: "Message updated",
	})

}

func (h *MessageHandler) DeleteMessage(c *fiber.Ctx) error {
	var request models.Message

	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	log.Println(request)
	slog.Debug(fmt.Sprintf("delete message endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.MessageBusiness.DeleteMessage(ctx, request); err != nil {
		slog.Debug(err.Error())
		if errors.Is(err, e.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(entities.Response{
				Error:   e.ErrNotFound.Error(),
				Content: nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: "Message deleted",
	})

}
