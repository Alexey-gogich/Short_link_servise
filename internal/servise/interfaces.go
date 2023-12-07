package servise

import (
	"context"
	"short_link_servise/internal/entity"
)

type LinksService interface {
	Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error)
	Create(ctx context.Context, links *entity.LinkCreate) error
}

type Servises struct {
	Links LinksService
}
