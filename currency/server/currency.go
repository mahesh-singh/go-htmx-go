package server

import (
	"context"
	"log"

	pb "github.com/mahesh-singh/go-htmx-go/currency/protos/currency/protos"
)

type Currency struct {
	l *log.Logger
	pb.UnimplementedCurrencyServer
}

func NewCurrency(l *log.Logger) *Currency {
	return &Currency{l: l}
}

func (c *Currency) GetRate(ctx context.Context, req *pb.RateRequest) (*pb.RateResponse, error) {
	c.l.Printf("Handle request for GetRate", "base", req.GetBase(), "dest", req.GetDestination())
	return &pb.RateResponse{Rate: .5}, nil
}
