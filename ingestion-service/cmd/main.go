package main

import (
	"fmt"
	"log"
	"os"

	"github.com/irhendra09/ingestion-service/internal/config"
	"github.com/irhendra09/ingestion-service/internal/repository"
	"github.com/irhendra09/ingestion-service/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	fmt.Println("Starting...")
	client, err := config.NewMongo()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))

	repo := &repository.MongoRepo{
		Collection: collection,
	}
	postService := &service.PostService{
		Repo: repo,
	}

	err = postService.SyncPosts()

	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS SYNC POSTS")

	//cronScheduler := scheduler.CronScheduler{
	//	PostService: postService,
	//}
	//
	//cronScheduler.Start()
}
