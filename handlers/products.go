package handlers

import (
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
	if r.Method == http.MethodGet {
		p.getProducts(w)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter) {
	p.l.Println("Handle GET Products")
	d := data.GetProducts()
	err := d.ToJSON(w)
	if err != nil {
		p.l.Fatal(err)
		http.Error(w, "cannot marshall json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle ADD Product")
	pr := &data.Product{}
	err := pr.FromJSON(r.Body)
	if err != nil {
		p.l.Fatal(err)
		http.Error(w, "cannot unmarshall json", http.StatusInternalServerError)
	}
	data.AddProduct(pr)
}
