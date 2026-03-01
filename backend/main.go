package main

// TODO: Use a framework (fiber, gin, echo, etc.)
import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Serve the frontend index.html file with the initial data injected to increase initial page load speed.
// After initial page is loaded, the frontend will fetch the data from the backend like a SPA
func homeHandler(w http.ResponseWriter, r *http.Request, staticDir string) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	jsonData, err := json.Marshal(map[string]interface{}{
		"stories": []interface{}{},
	})
	if err != nil {
		log.Printf("Failed to marshal data: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	indexPath := filepath.Join(staticDir, "index.html")
	body, err := os.ReadFile(indexPath)
	if err != nil {
		log.Printf("Frontend not built: %v — run 'cd frontend && npm run build'", err)
		http.Error(w, "Frontend not built: run 'cd frontend && npm run build'", http.StatusInternalServerError)
		return
	}

	html := string(body)
	injectScript := `<script>window.__INITIAL_DATA__ = ` + string(jsonData) + `;</script>`
	if idx := strings.Index(html, "<body"); idx >= 0 {
		end := strings.Index(html[idx:], ">") + idx + 1
		html = html[:end] + injectScript + html[end:]
	} else {
		html = injectScript + html
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func main() {
	port := getEnv("PORT", "8080")
	staticDir := getEnv("STATIC_DIR", "../frontend/dist")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, staticDir)
	})

	// Serve assets from the frontend dist directory
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
