package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Contests          *mongo.Collection
	Gamers            *mongo.Collection
	FrequencyAnalyses *mongo.Collection
	GameResults       *mongo.Collection
	AppSettings       *mongo.Collection
}

func NewConnectMongo(env *Env) *Collection {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s/",
		env.MONGO_ROOT_USER,
		env.MONGO_ROOT_PASSWORD,
		env.MONGO_PORT,
	)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌​ Erro ao ligar ao MongoDB.", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌​ Erro ao ligar ao MongoDB.", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	collection := client.Database(env.MONGO_DB_NAME)

	c := Collection{
		Contests:          collection.Collection(env.MONGO_COLLECTION_CONTESTS),
		Gamers:            collection.Collection(env.MONGO_COLLECTION_GAMERS),
		FrequencyAnalyses: collection.Collection(env.MONGO_COLLECTION_FREQUENCY_ANALYSES),
		GameResults:       collection.Collection(env.MONGO_COLLECTION_GAME_RESULTS),
		AppSettings:       collection.Collection(env.MONGO_COLLECTION_APP_SETTINGS),
	}

	return &c
}
