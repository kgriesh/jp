package app

import (
	"encoding/json"
	"jp/app/db"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func New(db db.DbService) http.Handler {

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/dinosaurs", getDinos(db))
		//r.Get("/{itemId}", GetProduct(db))
		//r.Post("/", CreateProduct(db))
	})
	return router
}

func getDinos(db db.DbService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dinos, err := db.GetDinos()
		if err != nil {
			render.Render(w, r, ServerError(err))
			return
		}
		err = respondwithJSON(w, 200, &dinos)
		if err != nil {
			render.Render(w, r, ServerError(err))
			return
		}
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	return err
}
