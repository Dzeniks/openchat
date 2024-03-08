package databaseService

import (
	"context"
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
