package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InintMongoDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	conn := os.Getenv("MONGO_CONN")
	if conn == "" {
		log.Fatalf("MONGO_CONN not set in .env")
	}

	timeout := os.Getenv("DB_TIMEOUT")
	if timeout == "" {
		log.Fatalf("DB_TIMEOUT not set in .env")
	}

	if err != nil {
		log.Fatalf("Failed to convert env timeout.")
	}

	clientOpt := options.Client().ApplyURI(conn)

	to, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("Failed to convert time duration: %v\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), to)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		clientOpt,
	)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Can't connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}
