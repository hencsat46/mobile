package userhttphandler

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
	"hackathon/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService   IBusinessUser
	jwtMiddleware *jwt.JWT
}

type IBusinessUser interface {
	FetchUserChatrooms(ctx context.Context, GUID string) ([]models.Chatroom, error)
	Login(ctx context.Context, user models.User) (string, error)
	CreateUser(ctx context.Context, user models.User) (string, error)
	UpdateUsername(ctx context.Context, username, GUID string) error
	UpdateEmail(ctx context.Context, email, GUID string) error
	UpdatePassword(ctx context.Context, oldPassword, newPassword, GUID string) error
	DeleteUser(ctx context.Context, GUID string) error
	EnterChatroom(ctx context.Context, guid, cid string) error
	QuitChatroom(ctx context.Context, guid, cid string) error
}

func New(userService IBusinessUser, jwt *jwt.JWT) *UserHandler {
	return &UserHandler{
		UserService:   userService,
		jwtMiddleware: jwt,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var request entities.UserDTO

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("create user endpoint called: %v", request))

	user := models.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	guid, err := h.UserService.CreateUser(ctx, user)
	if err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   err.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(entities.Response{
		Error:   "",
		Content: guid,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request entities.UserDTO

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("login endpoint called: %v", request))

	userEntity := models.User{
		GUID:     request.GUID,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	guid, err := h.UserService.Login(ctx, userEntity)
	if err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   err.Error(),
			Content: nil,
		})
	}

	token := h.jwtMiddleware.CreateToken(guid)

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error: "",
		Content: struct {
			GUID  string `json:"guid"`
			Token string `json:"token"`
		}{
			GUID:  guid,
			Token: token,
		},
	})
}

func (h *UserHandler) UpdateUsername(c *fiber.Ctx) error {
	var request entities.UserDTO

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update username endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserService.UpdateUsername(ctx, request.Username, request.GUID); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(entities.Response{
		Error:   "",
		Content: "User updated",
	})
}

func (h *UserHandler) UpdateEmail(c *fiber.Ctx) error {
	var request entities.UserDTO

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update email endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserService.UpdateEmail(ctx, request.Email, request.GUID); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(entities.Response{
		Error:   "",
		Content: "Email updated",
	})
}

func (h UserHandler) UpdatePassword(c *fiber.Ctx) error {
	var request entities.UpdatePasswordDTO

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update password endpoint called: %v", request))

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserService.UpdatePassword(ctx, request.OldPassword, request.NewPassword, request.GUID); err != nil {
		slog.Debug(err.Error())
		if errors.Is(err, e.ErrPasswordIncorrect) {
			return c.Status(http.StatusBadRequest).JSON(entities.Response{
				Error:   e.ErrPasswordIncorrect.Error(),
				Content: nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(entities.Response{
		Error:   "",
		Content: "Password updated",
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	guid := c.Params("GUID")
	slog.Debug(fmt.Sprintf("delete user endpoint called: %v", guid))

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserService.DeleteUser(ctx, guid); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(entities.Response{
			Error:   e.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(entities.Response{
		Error:   "",
		Content: "User deleted",
	})
}

func (h *UserHandler) FetchUserChatrooms(c *fiber.Ctx) error {
	guid := c.Params("guid")

	if len(guid) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("fetch user's chatrooms endpoint called: %v", guid))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	chatrooms, err := h.UserService.FetchUserChatrooms(ctx, guid)
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
		Content: chatrooms,
	})
}

func (h *UserHandler) EnterChatroom(c *fiber.Ctx) error {
	cid := c.Params("cid")
	guid := c.Params("guid")
	slog.Debug(fmt.Sprintf("enter chatroom endpoint called %s", cid))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.UserService.EnterChatroom(ctx, guid, cid); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: "entered",
	})
}

func (h *UserHandler) ExitChatroom(c *fiber.Ctx) error {
	cid := c.Params("cid")
	guid := c.Params("guid")
	slog.Debug(fmt.Sprintf("exit chatroom endpoint called %s", cid))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.UserService.QuitChatroom(ctx, guid, cid); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(entities.Response{
			Error:   e.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(entities.Response{
		Error:   "",
		Content: "exited",
	})
}
