package scheduler

import (
	"log"

	"github.com/irhendra09/ingestion-service/internal/service"
	"github.com/robfig/cron/v3"
)

type CronScheduler struct {
	PostService *service.PostService
}

func (c *CronScheduler) Start() {

	scheduler := cron.New()

	_, err := scheduler.AddFunc("0 1 * * *", func() {

		log.Println("START SYNC POSTS")

		err := c.PostService.SyncPosts()

		if err != nil {
			log.Println("ERROR SYNC POSTS:", err)
			return
		}

		log.Println("SUCCESS SYNC POSTS")
	})

	if err != nil {
		log.Fatal(err)
	}

	scheduler.Start()

	log.Println("CRON SCHEDULER STARTED")

	select {}
}
