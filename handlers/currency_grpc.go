package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/singhpratik/microservice/grpc/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{log}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.Base, "destination", rr.Destination)
	return &protos.RateResponse{
		Rate: 0.5,
	}, nil
}
