package repository

import (
	"context"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepositoryInterface interface {
	CreateGame(game *model.Game) error
}

type GameRepository struct {
	collection *mongo.Collection
}

func NewGameRepository(game *mongo.Collection) *GameRepository {
	return &GameRepository{
		collection: game,
	}
}

func (r *GameRepository) CreateGame(game *model.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, game)
	if err != nil {
		return err
	}

	return nil
}
