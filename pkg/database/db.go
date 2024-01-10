package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dskydiver/ipfs-metadata-fetcher/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser = getEnvVariable("DATABASE_USER")
	dbPassword = getEnvVariable("DATABASE_PASSWORD")
	dbName = getEnvVariable("DATABASE_NAME")
	dbHost = getEnvVariable("DATABASE_HOST")
	dbPort = getEnvVariable("DATABASE_PORT")
}

func getEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func NewClient() *ent.Client {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
