package main

import (
	"fmt"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config"
	"github.com/maiconDeSouza/lotofacil_lab/internal/http/server"
)

func main() {
	env := config.InitENV()
	collention := config.NewConnectMongo(env)

	mux := server.InitServer(env)

	fmt.Println(collention, mux)
}
