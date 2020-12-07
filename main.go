package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"
	protos "github.com/singhpratik/microservice/grpc/currency"
	"github.com/singhpratik/microservice/handlers"
	"google.golang.org/grpc"
)

func getGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	conn := getGrpcClient("localhost:8090")
	cc := protos.NewCurrencyClient(conn)
	defer conn.Close()
	ph := handlers.NewProducts(l, cc)
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

	go func() {
		setUpGoServer()
	}()

	sig := <-sigChan
	l.Printf("Received signal %s to terminate. trying to shutdown gracefully", sig)
	tc, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

func setUpGoServer() error {
	hl := hclog.Default()
	gs := grpc.NewServer()
	cs := handlers.NewCurrency(hl)
	protos.RegisterCurrencyServer(gs, cs)
	ls, err := net.Listen("tcp", ":8090")
	if err != nil {
		hl.Info("Unable to get port 8090", err)
		return err
	}
	hl.Info("Running GRPC server")
	return gs.Serve(ls)
}
