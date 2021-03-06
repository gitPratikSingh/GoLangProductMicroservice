// Package classification Product API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/singhpratik/microservice/data"
	protos "github.com/singhpratik/microservice/grpc/currency"
)

type Products struct {
	l *log.Logger
	c protos.CurrencyClient
}

func NewProducts(l *log.Logger, c protos.CurrencyClient) *Products {
	return &Products{l, c}
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

// swagger:route GET /products prodcuts listProducts
// returns a list of products
// responses:
// 	200: ProductResponse

// getProduct returns the list of all products in the data store
func (p *Products) getProducts(w http.ResponseWriter) {
	p.l.Println("Handle GET Products")
	d := data.GetProducts()
	rate, err := p.c.GetRate(context.Background(), &protos.RateRequest{
		Base:        "USD",
		Destination: "GRP",
	})
	if err != nil {
		p.l.Fatal(err)
		http.Error(w, "cannot get exchange rate", http.StatusInternalServerError)
	}

	p.l.Printf("Rate of exchange %v", rate.Rate)

	for _, v := range d {
		v.Price = v.Price * rate.Rate
	}

	w.Header().Add("Content-Type", "application/json")
	err = d.ToJSON(w)
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
