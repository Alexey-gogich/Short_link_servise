package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/servise"

	pb "short_link_servise/internal/proto"

	"google.golang.org/grpc/status"
)

type linksHandler struct {
	serv servise.LinksService
}

func NewLinksHandler(serv servise.LinksService) *linksHandler {
	return &linksHandler{
		serv: serv,
	}
}

func (hand *linksHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	url, err := hand.serv.Get(ctx, &entity.ShortLink{
		Link: request.ShortUrl,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(5, "")
		}
		log.Println(fmt.Errorf("/servises: %w", err))
		return nil, status.Error(1, "Server error")
	}
	response := &pb.GetResponse{
		Url: url.Link,
	}
	return response, nil
}

func (hand *linksHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := hand.serv.Create(ctx, &entity.LinkCreate{
		Link: request.Url,
	})
	if err != nil {
		log.Println(fmt.Errorf("/servises: %w", err))
		return nil, status.Error(1, "Server error")
	}
	response := &pb.CreateResponse{}
	return response, nil
}
