package userservice

import (
	"context"
	"fmt"
	"hackathon/models"
	hash "hackathon/pkg/hasher"
	"log/slog"

	e "hackathon/exceptions"

	"github.com/beevik/guid"
)

type UserService struct {
	UserDao IDataAccessUser
}

type IDataAccessUser interface {
	FetchUserChatrooms(ctx context.Context, GUID string) ([]models.Chatroom, error)
	Login(ctx context.Context, user models.User) (string, error)
	CreateUser(ctx context.Context, user models.User) error
	UpdateUsername(ctx context.Context, username, GUID string) error
	UpdateEmail(ctx context.Context, email, GUID string) error
	UpdatePassword(ctx context.Context, newPassword, GUID string) error
	DeleteUser(ctx context.Context, GUID string) error
	GetUser(ctx context.Context, GUID string) (*models.User, error)
	GetUserByName(ctx context.Context, Username string) (*models.User, error)
	EnterChatroom(ctx context.Context, guid, cid string) error
	QuitChatroom(ctx context.Context, guid, cid string) error
}

func New(userdao IDataAccessUser) *UserService {
	return &UserService{
		UserDao: userdao,
	}
}

func (b *UserService) FetchUserChatrooms(ctx context.Context, GUID string) ([]models.Chatroom, error) {
	slog.Debug(fmt.Sprintf("fetching chatrooms: %v", GUID))

	chatrooms, err := b.UserDao.FetchUserChatrooms(ctx, GUID)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return chatrooms, nil
}

func (b *UserService) Login(ctx context.Context, user models.User) (string, error) {
	usr, err := b.UserDao.GetUserByName(ctx, user.Username)
	if err != nil {
		slog.Debug(err.Error())
		return "", e.ErrUserDoesNotExists
	}
	slog.Debug(fmt.Sprintf("%v SUKAAAAAAAA", usr))

	if !hash.Hshr.Validate(usr.Password, user.Password) {
		return "", e.ErrPasswordIncorrect
	}

	guid, err := b.UserDao.Login(ctx, user)
	if err != nil {
		slog.Debug(err.Error())
		return "", err
	}

	return guid, nil
}

func (b *UserService) CreateUser(ctx context.Context, user models.User) (string, error) {
	user.GUID = guid.NewString()

	user.Password = hash.Hshr.Hash(user.Password)

	usr, err := b.UserDao.GetUserByName(ctx, user.Username)
	if err != nil {
		slog.Debug(err.Error())
	}
	if usr != nil {
		return "", e.ErrSuchUserAlreadyExists
	}

	err = b.UserDao.CreateUser(ctx, user)
	if err != nil {
		slog.Debug(err.Error())
		return "", err
	}

	return user.GUID, nil
}

func (b *UserService) UpdateUsername(ctx context.Context, username, GUID string) error {
	if err := b.UserDao.UpdateUsername(ctx, username, GUID); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *UserService) UpdateEmail(ctx context.Context, email, GUID string) error {
	if err := b.UserDao.UpdateEmail(ctx, email, GUID); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

// TODO: logic with comparison of provided old password with real old password
func (b *UserService) UpdatePassword(ctx context.Context, oldPassword, newPassword, GUID string) error {
	user, err := b.UserDao.GetUser(ctx, GUID)
	slog.Debug(fmt.Sprintf("%v", user))
	if err != nil {
		slog.Debug(err.Error())
		return err
	}

	if !hash.Hshr.Validate(user.Password, oldPassword) {
		return e.ErrPasswordIncorrect
	}

	newPassword = hash.Hshr.Hash(newPassword)
	if err := b.UserDao.UpdatePassword(ctx, newPassword, GUID); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *UserService) DeleteUser(ctx context.Context, GUID string) error {
	if err := b.UserDao.DeleteUser(ctx, GUID); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *UserService) EnterChatroom(ctx context.Context, guid, cid string) error {
	if err := b.UserDao.EnterChatroom(ctx, guid, cid); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *UserService) QuitChatroom(ctx context.Context, guid, cid string) error {
	if err := b.UserDao.QuitChatroom(ctx, guid, cid); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
