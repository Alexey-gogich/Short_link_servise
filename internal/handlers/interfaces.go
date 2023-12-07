package handlers

import (
	"context"
	pb "short_link_servise/internal/proto"
)

type LinksHandler interface {
	Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error)
	Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error)
}

type Handlers struct {
	Links LinksHandler
}
