package repository

import (
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

func (repo *linksRepo) Get(link *entity.ShortLink) (*entity.Link, error) {
	result := entity.Link{}
	err := repo.db.QueryRowx("Select url from links where short_url = $1", link.Link).Scan(&result.Link)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *linksRepo) Create(links *entity.LinkCreate) error {
	_, err := repo.db.Exec("Insert into links (url, short_url) values ($1, $2)", links.Link, links.ShortLink)
	if err != nil {
		return err
	}
	return nil
}
