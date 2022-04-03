package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	hostname "github.com/naffandroo/go-backend-hostname/api/v1"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/heartbeat"))

	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	fmt.Println("🚏 Adding Routes...")
	r.Route("/v1", func(r chi.Router) {
		r.Get("/hostname", hostname.GetHostName)
	})

	fmt.Println("✅ Starting up...")
	http.ListenAndServe(":8080", r)
}
