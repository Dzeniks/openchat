package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server-go/lib/databaseService"
	env "server-go/lib/dotEnv"
)

func main() {

	env.LoadDotEnv()

	// Dot Env
	mongoUrl := env.DotEnv.MongoConnectionURI
	log.Print(mongoUrl)

	// Get Client
	client, err := databaseService.GetClient()
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, context.Background())
	_ = databaseService.GetDatabase(client)
	// Check if the user exists and is active

	log.Print("User does not exist")

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
