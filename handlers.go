package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ProcessReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := CalculatePoints(receipt)
	store.Lock()
	store.Receipts[id] = points
	store.Unlock()

	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	store.Lock()
	points, exists := store.Receipts[id]
	store.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
