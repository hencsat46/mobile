package chatroomHTTPhandler

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
	hubmanager "hackathon/internal/presentation/hubManager"
	"hackathon/models"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ChatroomHandler struct {
	*hubmanager.HubManager
	ChatroomBusiness IBusinessChatroom
}

type IBusinessChatroom interface {
	GetChatrooms(ctx context.Context) ([]models.Chatroom, error)
	CreateChatroom(ctx context.Context, chatroom models.Chatroom) (string, error)
	UpdateChatroom(ctx context.Context, chatroomID, chatroomName string) error
	DeleteChatroom(ctx context.Context, ownerGUID, chatroomID string) error
}

func New(hubmngr *hubmanager.HubManager, chatroomBusiness IBusinessChatroom) *ChatroomHandler {
	return &ChatroomHandler{
		HubManager:       hubmngr,
		ChatroomBusiness: chatroomBusiness,
	}
}

// GetChatroom godoc
// @Summary Получить все чаты
// @Description хз зачем, просто чтобы был список
// @Tags chatrooms
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} entities.Response{content=[]models.Chatroom} "Успешный ответ с массивом комнат"
// @Failure 400 {object} entities.Response "Bad request"
// @Router /chatroom/get [get]
func (h *ChatroomHandler) GetChatroom(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	chatrooms, err := h.ChatroomBusiness.GetChatrooms(ctx)
	if err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: chatrooms,
	})

}

// CreateChatroom godoc
// @Summary Создать чат
// @Description Создание чата
// @Tags chatrooms
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param chatroom body entities.ChatroomDTO true "Chatroom Data"
// @Success 201 {object} entities.Response "Chatroom created successfully"
// @Failure 400 {object} entities.Response "Bad request"
// @Router /chatroom/create [post]
func (h *ChatroomHandler) CreateChatroom(c *fiber.Ctx) error {
	var request entities.ChatroomDTO

	log.Println(string(c.Request().Body()))

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("create chatroom endpoint called: %v", request))

	chatroom := models.Chatroom{
		Name:      request.Name,
		OwnerGUID: request.OwnerGUID,
		IsPrivate: request.IsPrivate,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	createdID, err := h.ChatroomBusiness.CreateChatroom(ctx, chatroom)
	if err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	h.HubManager.LoadChatroomToHub(&entities.WSRoom{
		CID:          createdID,
		Participants: make(map[string]*websocket.Conn),
	})

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: createdID,
	})
}

// UpdateChatroom godoc
// @Summary Обновить чат
// @Description Обновление чата
// @Tags chatrooms
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param chatroom body entities.ChatroomDTO true "Chatroom Data"
// @Success 200 {object} entities.Response "Chatroom updated successfully"
// @Failure 400 {object} entities.Response "Bad request"
// @Router /chatroom/ [put]
func (h *ChatroomHandler) UpdateChatroom(c *fiber.Ctx) error {
	var request entities.ChatroomDTO

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update chatroom endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomBusiness.UpdateChatroom(ctx, request.ID, request.Name); err != nil {
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
		Content: "Chatroom updated",
	})
}

// DeleteChatroom godoc
// @Summary Удалить чат
// @Description Удаление чата
// @Tags chatrooms
// @Param Authorization header string true "Bearer token"
// @Param guid query string true "GUID пользователя"
// @Param cid query string true "ID чата"
// @Success 200 {object} entities.Response "Chatroom deleted successfully"
// @Failure 400 {object} entities.Response "Bad request"
// @Router /chatroom/{chatroomID}/{GUID} [delete]
func (h *ChatroomHandler) DeleteChatroom(c *fiber.Ctx) error {
	chatroomID := c.Params("chatroomID")
	ownerGUID := c.Params("GUID")
	slog.Debug(fmt.Sprintf("delete chatroom endpoint called: %v", chatroomID))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomBusiness.DeleteChatroom(ctx, ownerGUID, chatroomID); err != nil {
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
		Content: "Chatroom deleted",
	})
}
