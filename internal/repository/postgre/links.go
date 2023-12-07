package postgre

import (
	"context"
	"short_link_servise/internal/entity"

	"github.com/jmoiron/sqlx"
)

type linksRepo struct {
	db *sqlx.DB
}

func NewLinksRepo(db *sqlx.DB) *linksRepo {
	return &linksRepo{
		db: db,
	}
}

func (repo *linksRepo) Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error) {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	result := entity.Link{}
	err := repo.db.QueryRowxContext(dbCtx, "Select url from links where short_url = $1", link.Link).Scan(&result.Link)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *linksRepo) Create(ctx context.Context, links *entity.LinkCreate) error {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	_, err := repo.db.ExecContext(dbCtx, "Insert into links (url, short_url) values ($1, $2)", links.Link, links.ShortLink)
	if err != nil {
		return err
	}
	return nil
}
