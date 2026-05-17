package service

import (
	"log"

	"github.com/irhendra09/ingestion-service/internal/client"
	"github.com/irhendra09/ingestion-service/internal/produce"
	"github.com/irhendra09/ingestion-service/internal/repository"
	"github.com/irhendra09/ingestion-service/internal/transformer"
)

type PostService struct {
	Repo *repository.MongoRepo
}

func (s *PostService) SyncPosts() {

	response, err := client.GetPost()

	if err != nil {
		log.Println(err)
	}

	for _, item := range response.Post {

		post := transformer.ToPost(
			item.ID,
			item.Title,
			item.Body,
			item.Tags,
			item.Views,
		)

		err := s.Repo.Upsert(post)
		produce.Produce(&post)

		if err != nil {
			log.Println(err)
		}
	}
}
