package models

import "github.com/gofiber/contrib/websocket"

type Room struct {
	CID          string
	Participants map[string]*websocket.Conn
}

type User struct {
	GUID     string
	Username string
	Password string
	Email    string
}

type Message struct {
	MessageId  string `bson:"message_id" json:"message_id"`
	ChatroomId string `bson:"chatroom_id" json:"chatroom_id"`
	SenderGUID string `bson:"sender_guid" json:"sender_guid"`
	SenderName string `bson:"sender_name"`
	Content    string `bson:"content"`
	Image      bool   `bson:"image"`
}

type Chatroom struct {
	ChatroomId string `bson:"chatroom_id"`
	Name       string `bson:"name"`
	OwnerGUID  string `bson:"owner"`
	IsPrivate  bool   `bson:"isPrivate"`
}
