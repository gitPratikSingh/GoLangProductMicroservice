package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/singhpratik/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := data.GetProducts()
	j, err := json.Marshal(d)
	if err != nil {
		p.l.Fatal(err)
		http.Error(w, "cannot marshall json", http.StatusInternalServerError)
	}
	w.Write(j)
}
