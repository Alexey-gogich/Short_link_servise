package inmemory

import (
	"context"
	"short_link_servise/internal/entity"
)

type linksRepo struct {
	db Cache
}

func NewLinksRepo(db Cache) *linksRepo {
	return &linksRepo{
		db: db,
	}
}

func (repo *linksRepo) Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error) {
	var err error
	result := entity.Link{}
	result.Link, err = repo.db.Get(link)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *linksRepo) Create(ctx context.Context, links *entity.LinkCreate) error {
	err := repo.db.Insert(links)
	if err != nil {
		return err
	}
	return nil
}
