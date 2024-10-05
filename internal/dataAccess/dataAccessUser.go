package dataaccess

import (
	"context"
	"fmt"
	"hackathon/migrations"
	"hackathon/models"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DataAccess) FetchUserChatrooms(ctx context.Context, GUID string) ([]models.Chatroom, error) {
	var chatrooms []models.Chatroom

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")
	log.Println(GUID)
	filter := bson.D{{"owner", GUID}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	for cursor.Next(ctx) {
		result := models.Chatroom{}
		if err = cursor.Decode(&result); err != nil {
			slog.Debug(err.Error())
			return nil, err
		}
		chatrooms = append(chatrooms, result)

	}

	return chatrooms, nil
}

func (dao *DataAccess) Login(ctx context.Context, userData models.User) (string, error) {
	slog.Debug(fmt.Sprintf("login user %v", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"username", userData.Username}}

	mongoData := new(migrations.MongoUser)

	err := coll.FindOne(ctx, filter).Decode(&mongoData)
	if err != nil {
		slog.Debug(err.Error())
		return "", err
	}

	return mongoData.GUID, nil
}

func (dao *DataAccess) CreateUser(ctx context.Context, userData models.User) error {

	slog.Debug(fmt.Sprintf("creating user %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	data := migrations.MongoUser{
		GUID:           userData.GUID,
		Username:       userData.Username,
		HashedPassword: userData.Password,
		Email:          userData.Email,
		Chatrooms:      primitive.A{},
	}

	_, err := coll.InsertOne(ctx, data)

	if err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) UpdateUsername(ctx context.Context, newUsername, GUID string) error {
	slog.Debug(fmt.Sprintf("updating username %v", GUID))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", GUID}}

	update := bson.D{{"$set", bson.D{{"username", newUsername}}}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) UpdateEmail(ctx context.Context, newEmail, GUID string) error {
	slog.Debug(fmt.Sprintf("updating email %v", GUID))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", GUID}}

	update := bson.D{{"$set", bson.D{{"email", newEmail}}}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) UpdatePassword(ctx context.Context, newPassword, GUID string) error {
	slog.Debug(fmt.Sprintf("updating password %v", GUID))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", GUID}}

	update := bson.D{{"$set", bson.D{{"password", newPassword}}}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) DeleteUser(ctx context.Context, GUID string) error {
	slog.Debug(fmt.Sprintf("deleting user %v", GUID))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", GUID}}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) GetUserByName(ctx context.Context, Username string) (*models.User, error) {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"username": Username}

	result := &models.User{}

	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	slog.Debug(fmt.Sprintf("%v", result))
	return result, nil
}
