package config

import (
	"context"
	"fmt"
	utils "github.com/alicobanserver/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect() {
	//Database config :
	fmt.Println("connect is start")

	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.qz9j1.mongodb.net/<dbname>?retryWrites=true&w=majority")

	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
		log.Println("Server running port 3000")
	}
	db := client.Database("deneme")
	utils.UserCollection(db)
	return
}
