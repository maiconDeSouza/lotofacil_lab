package repository

import (
	"context"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppSettingsRepositoryInterface interface {
	FindAppSettings(app *model.AppSettings) error
}

type AppSettingsRepository struct {
	collection *mongo.Collection
}

func NewAppSettingsRepository(appSettings *mongo.Collection) *AppSettingsRepository {
	return &AppSettingsRepository{
		collection: appSettings,
	}
}

func (r *AppSettingsRepository) FindAppSettings(app *model.AppSettings) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	if err := r.collection.FindOne(ctx, bson.M{}).Decode(&app); err != nil {
		return err
	}

	return nil
}
