package dataaccess

import (
	"context"
	"fmt"
	"hackathon/migrations"
	"hackathon/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DataAccess) EnterChatroom(ctx context.Context, guid, cid string) error {
	slog.Debug(fmt.Sprintf("entering chatroom with guid: %v, cid: %v", guid, cid))

	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"guid": guid}

	update := bson.M{"$addToSet": bson.M{"chatrooms": cid}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) QuitChatroom(ctx context.Context, guid, cid string) error {
	slog.Debug(fmt.Sprintf("quitting chatroom with guid: %v, cid: %v", guid, cid))

	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"guid": guid}

	update := bson.M{"$pull": bson.M{"chatrooms": cid}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) CreateChatroom(ctx context.Context, chatroom models.Chatroom) error {
	slog.Debug(fmt.Sprintf("creating chatroom %v", chatroom))

	mongoChatroom := migrations.MongoChatroom{
		ChatroomId: chatroom.ChatroomId,
		Name:       chatroom.Name,
		OwnerGUID:  chatroom.OwnerGUID,
		IsPrivate:  chatroom.IsPrivate,
		Messages:   primitive.A{},
	}

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	if _, err := coll.InsertOne(ctx, mongoChatroom); err != nil {
		slog.Debug(err.Error())
		return err
	}

	coll = dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"guid": chatroom.OwnerGUID}

	update := bson.M{"$push": bson.M{"chatrooms": chatroom.ChatroomId}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) UpdateChatroom(ctx context.Context, chatroomID, chatroomName string) error {
	slog.Debug(fmt.Sprintf("updating chatroom %v", chatroomID))

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomID}

	update := bson.M{"$set": bson.M{"name": chatroomName}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) DeleteChatroom(ctx context.Context, ownerGUID, chatroomID string) error {
	slog.Debug(fmt.Sprintf("deleting chatroom %v", chatroomID))
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomID}

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		slog.Debug(err.Error())
		return err
	}

	coll = dao.mongoConnection.Database("ringo").Collection("users")
	filter = bson.M{"guid": ownerGUID}
	update := bson.M{"$pull": bson.M{"chatrooms": chatroomID}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) GetChatrooms(ctx context.Context) ([]models.Chatroom, error) {
	slog.Debug("getting chatrooms")
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")
	var chats []models.Chatroom

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	for cursor.Next(ctx) {
		var chat models.Chatroom
		if err := cursor.Decode(&chat); err != nil {
			slog.Debug(err.Error())
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}
