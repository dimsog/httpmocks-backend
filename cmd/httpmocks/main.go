package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middlewareLogger "github.com/dimsog/httpmocks-backend/internal/http-server/middleware/logger"
	"github.com/dimsog/httpmocks-backend/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log := logger.New()
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middlewareLogger.New(log))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	go func() {
		log.Debug("Server starting...")
		err := srv.ListenAndServe()

		if err != nil {
			log.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Debug("Server shutdown")
}
