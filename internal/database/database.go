package database

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *mongo.Client
}

var (
	// host = os.Getenv("DB_HOST")
	// port = os.Getenv("DB_PORT")
	// dbUri = os.Getenv("DB_URI")
	//database = os.Getenv("DB_DATABASE")
)

func NewConnectToDB(dbUri string) (*service, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbUri))

	if err != nil {
		log.Fatal(err)
		return nil, err;
	}

	return &service{
		db: client,
	}, nil;
}

func (s *service) Health() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
		return nil, err;
	}

	return map[string]string{
		"message": "It's healthy",
	}, nil
}
