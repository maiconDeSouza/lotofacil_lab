package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
	"github.com/maiconDeSouza/lotofacil_lab/internal/http/handler"
)

func InitServer(env *env.Env, handler handler.ContestHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /start", handler.Start)

	port := fmt.Sprintf("%s", env.SERVER_PORT)
	fmt.Println("🚀 Servidor iniciado com sucesso na porta: ", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
	return mux
}
