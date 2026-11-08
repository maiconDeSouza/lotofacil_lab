package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maiconDeSouza/lotofacil_lab/internal/config/env"
)

func InitServer(env *env.Env) *http.ServeMux {
	mux := http.NewServeMux()

	port := fmt.Sprintf("%s", env.SERVER_PORT)
	fmt.Println("🚀 Servidor iniciado com sucesso na porta: ", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
	return mux
}
