package server

import (
	"fmt"
	"log"
	"net/http"
)

type config struct {
	fileServerHits int
}

func Run() {
	host := "localhost"
	port := 8080

	mux := http.NewServeMux()
	chirpyMux := middlewareCors(mux)

	config := config{
		fileServerHits: 0,
	}

	// mux.HandleFunc("/", response404)
	mux.HandleFunc("/healthz", responseHealthz)
	mux.HandleFunc("/metrics", config.responseMetrics)
	mux.HandleFunc("/reset", config.responseReset)
	mux.Handle("/app/", http.StripPrefix("/app", config.middlewareHitCounter(http.FileServer(http.Dir("./public")))))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), chirpyMux))
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
