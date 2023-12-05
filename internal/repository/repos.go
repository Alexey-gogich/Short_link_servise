package repository

import (
	"short_link_servise/internal/entity"

	"github.com/jmoiron/sqlx"
)

type LinksRepo interface {
	Get(link *entity.ShortLink) (*entity.Link, error)
	Create(links *entity.LinkCreate) error
}

type Repos struct {
	Links LinksRepo
}

func NewRepos(db *sqlx.DB) *Repos {
	return &Repos{
		Links: NewLinksRepo(db),
	}
}
