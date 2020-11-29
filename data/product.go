package data

import (
	"encoding/json"
	"io"
	"time"
)

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

type Products []*Product

func GetProducts() Products {
	return ProductList
}

func (p *Products) ToJSON(w io.Writer) error {
	en := json.NewEncoder(w)
	return en.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	de := json.NewDecoder(r)
	return de.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}
