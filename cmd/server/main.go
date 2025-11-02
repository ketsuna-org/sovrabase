package main

import (
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
	"github.com/ketsuna-org/sovrabase/internal/config"
	"github.com/ketsuna-org/sovrabase/internal/middleware"
)

func main() {
	// Get config path from environment variable or use default
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load config from %s: %v", configPath, err)
	}

	// Setup HTTP Server

	router := mux.NewRouter()
	log.Printf("Starting API server on %s", cfg.API.APIAddr)
	if cfg.API.Domain != "" {
		log.Printf("  - Configured domain: %s", cfg.API.Domain)
	}
	if len(cfg.API.CORSAllow) > 0 {
		log.Printf("  - CORS allowed origins: %v", cfg.API.CORSAllow)
	}

	// Appliquer le middleware CORS
	corsConfig := &middleware.CORSConfig{
		Domain:         cfg.API.Domain,
		AllowedOrigins: cfg.API.CORSAllow,
	}
	router.Use(middleware.CORSMiddleware(corsConfig))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sovrabase API is running"))
	})

	if err := http.ListenAndServe(cfg.API.APIAddr, router); err != nil {
		log.Fatalf("failed to start API server: %v", err)
	}
}
