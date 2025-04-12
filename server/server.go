package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func StartServer(db *sql.DB) error {
	mux := http.NewServeMux()
	handler := KubernetesPod(mux)
	addBaseRoutes(mux, db)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	return err
}

func addBaseRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/", LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Root endpoint
		err := db.Ping()
		if err != nil {
			http.Error(w, `{"error": "Database connection failed"}`, http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"message": "Welcome to Jubilant-Engine!",
			"image":   "peterjbishop/jubilant-engine:latest",
			"volume":  "postgres",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}))
	mux.HandleFunc("/health", LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Health check endpoint
		response := map[string]interface{}{
			"message": "Welcome to Jubilant-Engine Status Check!",
			"image":   "peterjbishop/jubilant-engine:latest",
			"volume":  "postgres"}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}))
}
