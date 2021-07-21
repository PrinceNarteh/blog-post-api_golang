package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Ctx    context.Context
	DB     *mongo.Database
	Post   *mongo.Collection
	User   *mongo.Collection
)

func Setup() {
	// creating context
	var cancel context.CancelFunc
	Ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error

	// creating client
	clientOptions := options.Client().ApplyURI("mongodb://localhost")
	Client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	err = Client.Connect(Ctx)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}
	defer Client.Disconnect(Ctx)

	DB = Client.Database("blog")
	Post = DB.Collection("posts")
	User = DB.Collection("users")
}
