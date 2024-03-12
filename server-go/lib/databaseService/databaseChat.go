package databaseService

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMessageToChat(message Message, chatId string, database *mongo.Database) error {
	filter := bson.M{"_id": chatId}
	update := bson.M{"$push": bson.M{"messages": message}}
	_, err := database.Collection("chats").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetChatByID(chatID string, database *mongo.Database) (*Chat, error) {
	filter := bson.M{"_id": chatID}
	var chat Chat
	err := database.Collection("chats").FindOne(context.Background(), filter).Decode(&chat)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func CreateChat(ownerID string, userIDs []string, database *mongo.Database) (*string, error) {
	// Create mongo db chat add _id
	ID := uuid.New()

	// Combine ownerID with ID
	chatID := ownerID + "_" + ID.String()

	newChat := Chat{
		ChatID:   chatID,
		OwnerID:  ownerID,
		Users:    userIDs,
		Messages: []Message{},
	}

	res, err := database.Collection("chats").InsertOne(context.Background(), newChat)
	if err != nil {
		return nil, err
	}
	ChatID := res.InsertedID.(string)
	return &ChatID, nil
}
