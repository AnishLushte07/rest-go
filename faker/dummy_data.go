package faker

import (
	"fmt"
	"rest-go/entities"
)

var Products []entities.Product

func init() {
	AddProduct(&entities.Product{ID: 1, Name: "Anish Lushte", Price: 3.14, Description: "Hello world"})
	AddProduct(&entities.Product{ID: 2, Name: "Anish Lushte 2", Price: 3.1122, Description: "Hello world 2"})
}

func AddProduct(p *entities.Product) *[]entities.Product {
	Products = append(Products, *p)
	return &Products
}

func DeleteProduct(id uint) error {
	for i, t := range Products {
		if t.ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Product Id not found")
}
