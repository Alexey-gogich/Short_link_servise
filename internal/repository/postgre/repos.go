package postgre

import (
	"short_link_servise/internal/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	QueryTimeout = 10 * time.Second
)

func NewRepos(db *sqlx.DB) *repository.Repos {
	return &repository.Repos{
		Links: NewLinksRepo(db),
	}
}
