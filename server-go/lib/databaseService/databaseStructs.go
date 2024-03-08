package databaseService

import (
	"time"
)

type User struct {
	UserID           string    `bson:"_id"`
	Username         string    `bson:"username"`
	Email            string    `bson:"email"`
	Password         string    `bson:"password"`
	RegistrationDate time.Time `bson:"registration_date"`
	LastLoginDate    time.Time `bson:"last_login_date"`
	Active           bool      `bson:"active"`
	RefreshToken     string    `bson:"refresh_token"`
}

type Message struct {
	MessageID string    `bson:"_id"`
	SenderID  string    `bson:"sender_id"`
	Content   string    `bson:"content"`
	SentAt    time.Time `bson:"sent_at"`
}

type Chat struct {
	ChatID   string    `bson:"_id"`
	Users    []string  `bson:"users"`
	Messages []Message `bson:"messages"`
}
