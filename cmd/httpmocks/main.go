package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	srv.ListenAndServe()
}
