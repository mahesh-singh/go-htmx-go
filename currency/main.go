package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/mahesh-singh/go-htmx-go/currency/protos/currency/protos"
	"github.com/mahesh-singh/go-htmx-go/currency/server"
)

func main() {
	grpcServer := grpc.NewServer()

	l := log.New(os.Stdout, "currency-srv", log.LstdFlags)

	currencySrv := server.NewCurrency(l)

	pb.RegisterCurrencyServer(grpcServer, currencySrv)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	ls, err := net.Listen("tcp", ":9092")

	if err != nil {
		l.Fatal("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	grpcServer.Serve(ls)
}
