package server

import (
	"encoding/json"
	"log"
	"net/http"

	db "database-svc/db/sqlc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Router(store *db.Store) http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	log.Println("Pardon us while we list this store:", store)

	mux.Get("/party", partyHard)

	return mux
}

func partyHard(w http.ResponseWriter, _ *http.Request){
	payload := jsonResponse{
		Error: false,
		Message: "Welcome to the database jungle!",
	}

	writeJSON(w, http.StatusAccepted, payload)
}

func listAccounts(store *db.Store, w http.ResponseWriter, _ *http.Request){
	payload := jsonResponse{
		Error: false,
		Message: "Welcome to the database jungle!",
	}

	writeJSON(w, http.StatusAccepted, payload)
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}