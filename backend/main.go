package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/signalflowsean/news/backend/internal/neo4j"
)

type PageData struct {
	Title       string
	InitialData template.JS
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8080")
	staticDir := getEnv("STATIC_DIR", "../frontend/dist")
	neo4jURI := getEnv("NEO4J_URI", "neo4j://localhost:7687")
	neo4jUser := getEnv("NEO4J_USER", "neo4j")
	neo4jPassword := os.Getenv("NEO4J_PASSWORD")

	if neo4jPassword == "" {
		log.Fatal("NEO4J_PASSWORD environment variable is required")
	}

	ctx := context.Background()

	client, err := neo4j.NewClient(ctx, neo4jURI, neo4jUser, neo4jPassword)
	if err != nil {
		log.Fatalf("Failed to connect to Neo4j: %v", err)
	}
	defer client.Close(ctx)

	tmpl, err := template.ParseFiles("internal/templates/layout.gohtml")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		stories, err := client.GetStories(ctx)
		if err != nil {
			log.Printf("Failed to fetch stories: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(map[string]interface{}{
			"stories": stories,
		})
		if err != nil {
			log.Printf("Failed to marshal data: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:       "News Graph",
			InitialData: template.JS(jsonData),
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Failed to execute template: %v", err)
		}
	})

	assetsDir := filepath.Join(staticDir, "assets")
	fs := http.FileServer(http.Dir(assetsDir))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server starting on http://localhost:%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
