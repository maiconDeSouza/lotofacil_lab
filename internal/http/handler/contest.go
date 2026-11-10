package handler

import (
	"encoding/json"
	"net/http"

	"github.com/maiconDeSouza/lotofacil_lab/internal/service"
)

type ContestHandlerInterface interface {
}

type ContestHandler struct {
	contestService service.ContestServiceInterface
}

func NewContestHandler(contestService service.ContestServiceInterface) *ContestHandler {
	return &ContestHandler{
		contestService: contestService,
	}
}

func (h *ContestHandler) Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := h.contestService.Start()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
