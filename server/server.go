package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type config struct {
	fileServerHits int
}

func Run() {
	host := "localhost"
	port := 8080

	config := config{
		fileServerHits: 0,
	}

	adminRouter := chi.NewRouter()
	adminRouter.Use(middlewareCors)
	adminRouter.Get("/metrics", config.responseMetrics)

	apiRouter := chi.NewRouter()
	apiRouter.Use(middlewareCors)
	apiRouter.Get("/healthz", responseHealthz)
	apiRouter.Get("/reset", config.responseReset)

	r := chi.NewRouter()
	r.Use(middlewareCors)
	r.Handle("/app/*", http.StripPrefix("/app", config.middlewareHitCounter(http.FileServer(http.Dir("./public")))))
	r.Handle("/app", http.StripPrefix("/app", config.middlewareHitCounter(http.FileServer(http.Dir("./public")))))
	r.Mount("/admin", adminRouter)
	r.Mount("/api", apiRouter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r))
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (c *config) middlewareHitCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.fileServerHits++
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}
