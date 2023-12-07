package repository

import (
	"context"
	"short_link_servise/internal/entity"
)

type LinksRepo interface {
	Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error)
	Create(ctx context.Context, links *entity.LinkCreate) error
}

type Repos struct {
	Links LinksRepo
}
