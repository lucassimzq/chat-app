package models

type Message struct {
	Message string 	`bson:"message,omitempty"`
	ChatID int 	`bson:"chat_id"`
	UserID int 	`bson:"user_id"`
}