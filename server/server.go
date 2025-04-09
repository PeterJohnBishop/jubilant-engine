package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StartServer() error {
	mux := http.NewServeMux()
	handler := KubernetesPod(mux)
	addBaseRoutes(mux)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	return err
}

func addBaseRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"message":      "Welcome to Jubilant-Engine!",
			"docker image": "peterjbaker/jubilant-engine:latest",
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
			"message":      "Welcome to Jubilant-Engine!",
			"docker image": "peterjbaker/jubilant-engine:latest",
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
}
