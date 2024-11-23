package db

import (
	"context"
	"log"

	"github.com/G-Villarinho/parallelizing-golang-ops/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, *mongo.Collection) {
	clientOptions := options.Client().ApplyURI(config.MongoDBConfig())
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	collection := client.Database("school").Collection("students")
	return client, collection
}
