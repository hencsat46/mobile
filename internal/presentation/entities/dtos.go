package entities

type ChatroomDTO struct {
	ID                string `json:"chatroom_id"`
	Name              string `json:"name"`
	OwnerGUID         string `json:"guid"`
	IsPrivate         bool   `json:"is_private"`
	ParticipantsLimit int    `json:"participants_limit"`
}

type CreateChatroom struct {
	Guid string `json:"guid"`
	Name string `json:"name"`
	IsPrivate bool `json:"is_private"`
}

type UserDTO struct {
	GUID     string `json:"guid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUsername struct {
	GUID string `json:"guid"`
	Username string `json:"username"`
}

type UpdateEmail struct {
	GUID string `json:"guid"`
	Email string `json:"email"`
}

type TokenResponse struct {
	Error string
	Content struct {
		Token string `json:"Token"`
		UserGuid string `json:"UserGuid"`
	}
}

type UpdatePasswordDTO struct {
	GUID        string `json:"guid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"password"`
}

type MessageDTO struct {
	MessageId  string `json:"message_id"`
	ChatroomID string `json:"chatroom_id"`
	Content    string `json:"content"`
}

type Response struct {
	Error   string `json:"error"`
	Content any    `json:"content"`
}
