package databaseService

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser Inset User into MongoDB database
func CreateUser(User *User, database *mongo.Database) error {

	_, err := database.Collection("users").InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(userID *string, database *mongo.Database) (*User, error) {
	var user User
	filter := bson.M{"_id": userID}
	err := database.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email *string, database *mongo.Database) (*User, error) {
	var user User
	filter := bson.M{"email": email}
	err := database.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertRefreshToken(user *User, database *mongo.Database) error {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"refresh_token": user.RefreshToken}}
	_, err := database.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLoginDate(user *User, database *mongo.Database) error {
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"last_login_date": user.LastLoginDate}}
	_, err := database.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
