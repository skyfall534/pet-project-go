package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ID string `json:"id"`
}

var products = []Product{{Name: "apple", Description: "its an apple", ID: "1"}}

func main () {
	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}) 
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		// var products = []string{"apple", "banana", "pineapple"}
		marshalProducts, err := json.Marshal(products)
		if err != nil {

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalProducts)
	})
	r.Put("/product/{id}", func (w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "id")
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {

		}
		for _, product := range products {
			if (product.ID == productId) {
				product = newProduct
			}
		}

		jsonProduct, err := json.Marshal(newProduct)
		if err != nil {

		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonProduct)
		
		
	})
	http.ListenAndServe(":8080", r)
}
