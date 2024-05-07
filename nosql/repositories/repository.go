package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type Repository struct {
	collection *mongo.Collection
}

func (s *Repository) Create(ctx context.Context, model interface{}) error {
	_, err := s.collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
func (s *Repository) Update(ctx context.Context, id string, values bson.M) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": values}
	op := options.Update().SetUpsert(false)

	_, err := s.collection.UpdateOne(ctx, filter, update, op)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) GetById(ctx context.Context, id string, model interface{}) error {
	filter := bson.M{"_id": id}

	result := s.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}

	modelType := reflect.TypeOf(model)
	if modelType.Kind() != reflect.Ptr {
		return errors.New("model must be a pointer")
	}

	err := result.Decode(model)
	if err != nil {
		return err
	}

	return nil
}
