package controller

import (
	"context"
	"errors"
	"os"
	ent "root/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectDatabase(log *logrus.Logger) (*ent.Client, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
		return nil, errors.New("DATABASE_URL is not set")
	}

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	log.Info("connected to the database successfully")

	autoMigrate(client, log)

	log.Info("successfully created schema resources")

	return client, nil
}

func autoMigrate(client *ent.Client, log *logrus.Logger) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
