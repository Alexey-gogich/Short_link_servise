package servise

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"short_link_servise/internal/entity"
	"short_link_servise/internal/repository"
	"short_link_servise/internal/utils"
)

type linkServise struct {
	repo  repository.LinksRepo
	utils utils.LinkUtils
}

func NewLinksServise(links repository.LinksRepo) *linkServise {
	return &linkServise{
		repo:  links,
		utils: utils.NewLinksUtils(),
	}
}

func (serv *linkServise) Get(ctx context.Context, link *entity.ShortLink) (*entity.Link, error) {
	short, err := serv.repo.Get(ctx, link)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, fmt.Errorf("/repository: %w", err)
		}
	}
	return short, nil
}

func (serv *linkServise) Create(ctx context.Context, links *entity.LinkCreate) error {
	links.ShortLink = serv.utils.GenerateShortUrl()
	err := serv.repo.Create(ctx, links)
	if err != nil {
		UniqueError, regex_err := regexp.MatchString(err.Error(), "повторяющееся значение")
		if regex_err != nil {
			return regex_err
		}
		if UniqueError {
			for {
				links.ShortLink = serv.utils.GenerateShortUrl()
				err := serv.repo.Create(ctx, links)
				if err == nil {
					break
				}
			}
		} else {
			return fmt.Errorf("/repository: %w", err)
		}
	}
	return nil
}
