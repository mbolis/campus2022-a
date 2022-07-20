package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mbolis/yello/db"
	"github.com/mbolis/yello/model"
)

func main() {
	routes := chi.NewRouter()
	routes.Use(middleware.Logger)
	routes.Get("/lists", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(db.GetAllLists())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(bytes)
	})
	routes.Post("/lists", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var l model.List
		err = json.Unmarshal(bytes, &l)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = db.CreateNewList(&l)
		if err != nil {
			http.Error(w, err.Error(), 409)
			return
		}
		bytes, err = json.Marshal(l)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(201)
		w.Write(bytes)
	})

	http.ListenAndServe(":5000", routes)
}
