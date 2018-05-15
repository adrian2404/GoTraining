package web

import (
	"GoTraining/storage"
	"github.com/go-chi/chi"
	"net/http"
	"github.com/pkg/errors"
	"log"
	"encoding/json"
)

type Storage interface {
	GetGiphs() (storage.Giphs, error)
	GetGiph(giphID string) (storage.Giph)
	Close() error
}

type handler struct {
	storage Storage
}

func NewHandler (storage Storage) *handler {
	return &handler{storage: storage}
}

func Server(h *handler, port string) error {
	r := chi.NewRouter()
	r.Get("/", h.giphsHandlerGet)
	return http.ListenAndServe(port, r)
}


func (h *handler) giphsHandlerGet(w http.ResponseWriter, r *http.Request){
	giphs, err := h.storage.GetGiphs()
	if err != nil {
		err = errors.Wrap(err, "Failed to fetch giphs")
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(giphs)
	if err != nil {
		log.Println(err)
		return
	}
}