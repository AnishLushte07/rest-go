package controllers

import (
	"encoding/json"
	"net/http"
	"rest-go/entities"
	"rest-go/faker"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product

	json.NewDecoder(r.Body).Decode(&product)
	data := faker.AddProduct(&product)

	json.NewEncoder(w).Encode(&data)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := faker.Products
	json.NewEncoder(w).Encode(&data)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	uProductId := uint(productId)

	error := faker.DeleteProduct(uProductId)

	if error != nil {
		// w.Write([]byte(error.Error()))
		http.Error(w, error.Error(), http.StatusBadRequest)
	}
}
