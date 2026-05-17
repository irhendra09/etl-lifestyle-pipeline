package repository

import (
	"context"

	"github.com/irhendra09/ingestion-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Collection *mongo.Collection
}

func (r MongoRepo) Upsert(post model.Post) error {
	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": post}
	opts := options.Update().SetUpsert(true)
	_, err := r.Collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}
