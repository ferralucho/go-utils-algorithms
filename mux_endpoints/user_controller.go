package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// Register you route here
	r.HandleFunc("/users", QueryHandler).Methods("Get")
	return r
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	userName := queryParams.Get("name")[0]

	users := []

	existingUsers := database.GetUsers

	for _, user := range existingUsers {
		if user.name == userName {
			users = append(users, user)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}


