package data

import "time"

type Product struct {
	Name        string  `json:"name,omitempty"`
	ID          int     `json:"id,omitempty"`
	SKU         string  `json:"sku,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	CreatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

var ProductList = []*Product{
	&Product{
		Name:        "Latte",
		ID:          1,
		Price:       1.00,
		Description: "frothy milky coffee",
		DeletedAt:   "",
		CreatedAt:   time.Now().UTC().String(),
	},
	&Product{
		Name:        "Expresso",
		ID:          2,
		Price:       1.50,
		Description: "Short and strong coffee without milk",
		DeletedAt:   "",
		CreatedAt:   time.Now().UTC().String(),
	},
}

func GetProducts() []*Product {
	return ProductList
}
