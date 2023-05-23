package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-go/controllers"
	"rest-go/db"

	"github.com/gorilla/mux"
)

func main() {
	LoadConfig()

	db.Connect(AppConfig.DBName)
	db.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Println(http.ListenAndServe(fmt.Sprintf("localhost:%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	// router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	// router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
