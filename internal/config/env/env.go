package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Lottery struct {
	Name          string
	MinNumber     string
	MaxNumber     string
	NumberPerGame string
}

type Env struct {
	SERVER_PORT string

	MONGO_DB_NAME                       string
	MONGO_COLLECTION_CONTESTS           string
	MONGO_COLLECTION_GAMERS             string
	MONGO_COLLECTION_NUMBER_FREQUENCY   string
	MONGO_COLLECTION_FREQUENCY_ANALYSES string
	MONGO_COLLECTION_GAME_RESULTS       string
	MONGO_COLLECTION_APP_SETTINGS       string
	MONGO_ROOT_USER                     string
	MONGO_ROOT_PASSWORD                 string
	MONGO_PORT                          string

	JWT_SECRET_KEY string

	Lottery Lottery
}

func InitENV() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	l := Lottery{
		Name:          os.Getenv("Name"),
		MinNumber:     os.Getenv("MinNumber"),
		MaxNumber:     os.Getenv("MaxNumber"),
		NumberPerGame: os.Getenv("NumberPerGame"),
	}

	e := Env{
		SERVER_PORT:                         os.Getenv("SERVER_PORT"),
		MONGO_DB_NAME:                       os.Getenv("MONGO_DB_NAME"),
		MONGO_COLLECTION_CONTESTS:           os.Getenv("MONGO_COLLECTION_CONTESTS"),
		MONGO_COLLECTION_GAMERS:             os.Getenv("MONGO_COLLECTION_GAMERS"),
		MONGO_COLLECTION_NUMBER_FREQUENCY:   os.Getenv("MONGO_COLLECTION_NUMBER_FREQUENCY"),
		MONGO_COLLECTION_FREQUENCY_ANALYSES: os.Getenv("MONGO_COLLECTION_FREQUENCY_ANALYSES"),
		MONGO_COLLECTION_GAME_RESULTS:       os.Getenv("MONGO_COLLECTION_GAME_RESULTS"),
		MONGO_COLLECTION_APP_SETTINGS:       os.Getenv("MONGO_COLLECTION_APP_SETTINGS"),
		MONGO_ROOT_USER:                     os.Getenv("MONGO_ROOT_USER"),
		MONGO_ROOT_PASSWORD:                 os.Getenv("MONGO_ROOT_PASSWORD"),
		MONGO_PORT:                          os.Getenv("MONGO_PORT"),
		JWT_SECRET_KEY:                      os.Getenv("JWT_SECRET_KEY"),
		Lottery:                             l,
	}

	return &e
}
