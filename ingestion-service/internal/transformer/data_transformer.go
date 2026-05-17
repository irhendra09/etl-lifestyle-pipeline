package transformer

import (
	"time"

	"github.com/irhendra09/ingestion-service/internal/model"
)

func ToPost(
	id int,
	title string,
	body string,
	tags []string,
	views int,
) model.Post {
	return model.Post{
		ID:          id,
		Title:       title,
		Body:        body,
		Tags:        tags,
		Views:       views,
		ProcessedAt: time.Now(),
	}
}
