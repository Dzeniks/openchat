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
