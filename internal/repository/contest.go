package repository

import (
	"context"
	"errors"
	"time"

	errapp "github.com/maiconDeSouza/lotofacil_lab/internal/config/errApp"
	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContestRepositoryInterface interface {
	FindContest(contestNumber int, contest *model.Contest) error
	LastContest(lastContest *model.Contest) error
	CreateContest(contest *model.Contest) error
	FindAllContest(contests *[]*model.Contest) error
}

type ContestRepository struct {
	collection *mongo.Collection
}

func NewContestRepository(contest *mongo.Collection) *ContestRepository {
	return &ContestRepository{
		collection: contest,
	}
}

func (r *ContestRepository) FindAllContest(contests *[]*model.Contest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &contests); err != nil {
		return err
	}

	return nil
}

func (r *ContestRepository) FindContest(contestNumber int, contest *model.Contest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"number": contest}
	err := r.collection.FindOne(ctx, filter).Decode(&contest)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrNoDocuments
		}
		return err
	}

	return nil
}

func (r *ContestRepository) LastContest(lastContest *model.Contest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	findOptions := options.FindOne().
		SetSort(bson.D{{Key: "number", Value: -1}})

	err := r.collection.FindOne(ctx, bson.M{}, findOptions).Decode(&lastContest)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContestRepository) CreateContest(contest *model.Contest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, contest)
	if err != nil {
		return err
	}

	return nil
}
