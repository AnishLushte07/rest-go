package controllers

import (
	"encoding/json"
	"net/http"
	"rest-go/db"
	"rest-go/entities"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	product := &entities.Product{}

	json.NewDecoder(r.Body).Decode(&product)
	// data := faker.AddProduct(&product)
	db.DbInstance.Create(&product)
	json.NewEncoder(w).Encode(&product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// data := faker.Products
	products := &[]entities.Product{}
	db.DbInstance.Find(&products)
	json.NewEncoder(w).Encode(&products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	uProductId := uint(productId)

	// error := faker.DeleteProduct(uProductId)
	// not checking for product missing condition
	var product entities.Product
	db.DbInstance.Delete(&product, uProductId)
}
