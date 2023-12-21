package databaseService

import (
	"context"
	"server-go/lib/dotEnv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient Return a MongoDB database client
func GetClient() (*mongo.Client, error) {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI(dotEnv.DotEnv.MongoConnectionURI).SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1", // Specify the mechanism here
		Username:      dotEnv.DotEnv.MongoUsername,
		Password:      dotEnv.DotEnv.MongoPassword,
	})
	// Connect to the MongoDB server
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetDatabase(client *mongo.Client) *mongo.Database {
	return client.Database(dotEnv.DotEnv.MongoDatabase)
}
