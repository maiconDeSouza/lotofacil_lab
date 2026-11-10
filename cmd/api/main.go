package main

import (
	"fmt"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/database"
	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/http/handler"
	"github.com/maiconDeSouza/lotofacil_lab/internal/http/server"
	"github.com/maiconDeSouza/lotofacil_lab/internal/repository"
	"github.com/maiconDeSouza/lotofacil_lab/internal/service"
)

func main() {
	env := env.InitENV()
	collention := database.NewConnectMongo(env)
	repoApp := repository.NewAppSettingsRepository(collention.AppSettings)
	repoContest := repository.NewContestRepository(collention.Contests)
	repoGame := repository.NewGameRepository(collention.Gamers)
	repoNumbersFrequency := repository.NewNumberFrequencyRepository(collention.FrequencyAnalyses)

	service := service.NewContestService(repoContest, repoApp, repoGame, repoNumbersFrequency, env)

	handler := handler.NewContestHandler(service)

	mux := server.InitServer(env, *handler)

	fmt.Println(collention, mux)
}
