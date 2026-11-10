package repository

import (
	"context"
	"time"

	errapp "github.com/maiconDeSouza/lotofacil_lab/internal/config/errApp"
	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NumberFrequencyInterface interface {
	FindAllNumberFrequency(numbers *[]*model.NumberFrequency) error
	UpdateNumberFrequency(n *model.NumberFrequency) error
}

type NumberFrequencyRepository struct {
	collection *mongo.Collection
}

func NewNumberFrequencyRepository(numberFrequency *mongo.Collection) *NumberFrequencyRepository {
	return &NumberFrequencyRepository{
		collection: numberFrequency,
	}
}

func (r *NumberFrequencyRepository) FindAllNumberFrequency(numbers *[]*model.NumberFrequency) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &numbers); err != nil {
		return err
	}

	return nil
}

func (r *NumberFrequencyRepository) UpdateNumberFrequency(n *model.NumberFrequency) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"_id": n.ID}

	re, err := r.collection.ReplaceOne(ctx, filter, n)
	if err != nil {
		return err
	}

	if re.MatchedCount == 0 {
		return errapp.ErrNoDocuments
	}

	return nil
}
