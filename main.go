package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	Name string `json:"name"`
}

func main () {
	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}) 
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		// var products = []string{"apple", "banana", "pineapple"}
		var products = Product{Name: "apple"}
		marshalProducts, err := json.Marshal(products)
		if err != nil {

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalProducts)
	})
	http.ListenAndServe(":8080", r)
}
