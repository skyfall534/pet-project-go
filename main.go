package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ID string `json:"id"`
}

var products = []Product{{Name: "apple", Description: "its an apple", ID: "1"}, {Name: "banana", Description: "its an banana", ID: "2"}, {Name: "orange", Description: "its an orange", ID: "3"}}

func main () {
	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}) 
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		conn, err := pgx.Connect(context.Background(), "postgres://gouser:gopassword@localhost:5432/godb");
		defer conn.Close(context.Background());

		rows, err := conn.Query(context.Background(), "SELECT * FROM products");
		products, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Product])

		marshalProducts, err := json.Marshal(products)
		if err != nil {

		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalProducts)
	})
	r.Get("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "id")

		conn, err := pgx.Connect(context.Background(), "postgres://gouser:gopassword@localhost:5432/godb");
		defer conn.Close(context.Background());

		rows, err := conn.Query(context.Background(), "select * from products where id=$1", productId);
		product, _ := pgx.CollectOneRow(rows, pgx.RowToStructByName[Product]);
		if (Product{}) == product {
				notFoundMsg, err := json.Marshal("Product not found")
				if (err != nil) {

				}
				w.WriteHeader(http.StatusNotFound)
				w.Write(notFoundMsg)
				return
			
		}

		currProductJson, err := json.Marshal(product)
		if (err != nil) {

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(currProductJson)
	})
	r.Put("/product/{id}", func (w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "id")
		var newProduct Product
		var newProducts []Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			
		}
		for _, product := range products {
			if (product.ID == productId) {
				newProducts = append(newProducts, newProduct)
			} else {
				newProducts = append(newProducts, product)
			}
		}
		products = newProducts
		jsonProduct, err := json.Marshal(newProduct)
		if err != nil {

		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonProduct)

	})
	http.ListenAndServe(":8080", r)
}
