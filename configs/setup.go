package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDb() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoUri()))

	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalln(err)
	}

	return client
}

var Db *mongo.Client = ConnectDb()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golangApi").Collection(collectionName)
}
