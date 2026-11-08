package main

import (
	"fmt"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/database"
	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/http/server"
)

func main() {
	env := env.InitENV()
	collention := database.NewConnectMongo(env)

	mux := server.InitServer(env)

	fmt.Println(collention, mux)
}
