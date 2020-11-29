package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/singhpratik/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/products", ph)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(s *http.Server) {
		fmt.Println("Listening for http calls")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}(s)

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Printf("Received signal %s to terminate. trying to shutdown gracefully", sig)
	tc, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
