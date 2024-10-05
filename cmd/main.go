package main

import (
	"context"
	"os/signal"

	chatroomservice "hackathon/internal/business/chatroomService"
	hubloaderservice "hackathon/internal/business/hubLoaderService"
	messageservice "hackathon/internal/business/messageService"
	userservice "hackathon/internal/business/userService"
	wsservice "hackathon/internal/business/wsService"
	dataaccess "hackathon/internal/dataAccess"
	handlers "hackathon/internal/presentation"
	"hackathon/pkg/config"
	"hackathon/pkg/jwt"
	"hackathon/pkg/logger"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), os.Interrupt)
	defer cancel()

	cfg := config.New()
	logger := logger.New(cfg)
	logger.SetAsDefault()

	mng, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Mongo))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dataaccess := dataaccess.NewDataAccess(mng)

	chatroomSvc := chatroomservice.New(dataaccess)
	hubLoaderSvc := hubloaderservice.New(dataaccess)
	messageSvc := messageservice.New(dataaccess)
	userSvc := userservice.New(dataaccess)
	wsSvc := wsservice.New(dataaccess)

	handler := handlers.NewHandler(cfg, fiber.New(), hubLoaderSvc, messageSvc, userSvc, wsSvc, chatroomSvc, jwt.New(cfg))

	go func() {
		if err := handler.Start(); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	if err := handler.Shutdown(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
