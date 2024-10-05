package dataaccess

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DataAccess) GetUser(ctx context.Context, GUID string) (*models.User, error) {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"guid": GUID}

	result := &models.User{}

	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	slog.Debug(fmt.Sprintf("%v", result))
	return result, nil
}

func (dao *DataAccess) GetChatroom(ctx context.Context, chatroomID string) (*models.Chatroom, error) {
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomID}

	result := &models.Chatroom{}

	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return result, nil
}
