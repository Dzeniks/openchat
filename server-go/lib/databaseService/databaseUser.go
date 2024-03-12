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
	// Also used for activity
	filter := bson.M{"_id": user.UserID}
	update := bson.M{"$set": bson.M{"last_login_date": user.LastLoginDate}}
	_, err := database.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetAllChats(UserID string, database *mongo.Database) (*[]Chat, error) {
	filter := bson.M{"owner_id": UserID}
	ctx := context.Background()
	var chats []Chat
	chats_cursor, err := database.Collection("chats").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer chats_cursor.Close(ctx)
	for chats_cursor.Next(ctx) {
		var chat Chat // Replace with your chat struct
		err := chats_cursor.Decode(&chat)
		if err == nil {
			chats = append(chats, chat)
		}
	}
	if err = chats_cursor.Err(); err != nil {
		return nil, err
	}
	return &chats, nil
}
