package repositories

import (
	"context"
	"github.com/google/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"personal-assistant/internal/config/constants/environments"
)

type Database struct {
	Client *mongo.Client
}

// NewDatabase creates a new Database
func NewDatabase() *Database {
	return &Database{}
}

// Connect to the MongoDB database
func (d *Database) Connect() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(environments.DatabaseUrl))
	if err != nil {
		logger.Fatalf("Error connecting to MongoDB: %v", err)
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatalf("Error pinging MongoDB: %v", err)
		return nil
	}

	d.Client = client

	return client.Database(environments.DatabaseName)
}

func (d *Database) ConnectWithUriAndDBName(uri string, dbName string) *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatalf("Error connecting to MongoDB: %v", err)
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatalf("Error pinging MongoDB: %v", err)
		return nil
	}
	d.Client = client

	return client.Database(dbName)
}

func (d *Database) Disconnect() {
	err := d.Client.Disconnect(context.Background())
	if err != nil {
		logger.Fatalf("Error disconnecting from MongoDB: %v", err)
	}
}
