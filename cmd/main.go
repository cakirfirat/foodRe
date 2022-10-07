package main

import (
	. "foodRe/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/nutritive", FoodHandler).Methods("POST")
	r.HandleFunc("/api/recipe", RecipeFood).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	log.Println("Server ending...")
}
