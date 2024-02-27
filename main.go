package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", ProcessReceiptsHandler)
	r.HandleFunc("/receipts/{id}/points", GetPointsHandler)
	return r
}

func main() {
	r := setupRouter()

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")

	http.ListenAndServe(":8080", nil)
}
