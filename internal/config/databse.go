package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
	"github.com/maiconDeSouza/lotofacil_lab/pkg/random"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Contests          *mongo.Collection
	Gamers            *mongo.Collection
	NumberFrequency   *mongo.Collection
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
		NumberFrequency:   collection.Collection(env.MONGO_COLLECTION_NUMBER_FREQUENCY),
		FrequencyAnalyses: collection.Collection(env.MONGO_COLLECTION_FREQUENCY_ANALYSES),
		GameResults:       collection.Collection(env.MONGO_COLLECTION_GAME_RESULTS),
		AppSettings:       collection.Collection(env.MONGO_COLLECTION_APP_SETTINGS),
	}

	count, err := c.AppSettings.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("❌ Erro ao verificar a coleção AppSettings:", err)
	}

	if count == 0 {
		log.Println("ℹ️ Criando seu jogo fixo. Salvando dados iniciais...")

		s := random.GenerateLotofacilNumbers15()
		f := model.AppSettings{
			ID:        env.Lottery.Name,
			FixedGame: s,
			CreatedAt: time.Now(),
		}

		if _, err := c.AppSettings.InsertOne(ctx, f); err != nil {
			log.Fatal("❌ Erro ao salvar dados iniciais em AppSettings:", err)
		}

		fmt.Println("✅ Seus números da sorte são: ", s)
	}

	count, err = c.NumberFrequency.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("❌ Erro ao verificar a coleção NumberFrequency:", err)
	}

	if count == 0 {
		log.Println("ℹ️ Criando lista do números mais sorteados. Salvando dados iniciais...")
		for i := range 25 {
			f := model.NumberFrequency{
				ID:     fmt.Sprintf("%d", i+1),
				Number: i + 1,
				Count:  0,
			}
			if _, err := c.NumberFrequency.InsertOne(ctx, f); err != nil {
				log.Fatal("❌ Erro ao salvar dados iniciais em NumberFrequency:", err)
			}
		}

		fmt.Println("✅ Números frquentes salvos com suceso: ")
	}

	return &c
}
