package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/uandersonricardo/jobbing-jobs-service/routes"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DATABASE"), options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

func main() {
	routes.InitRouter()
}
