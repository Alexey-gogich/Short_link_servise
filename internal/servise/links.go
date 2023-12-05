package servise

import (
	"fmt"
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

func (serv *linkServise) Get(link *entity.ShortLink) (*entity.Link, error) {
	short, err := serv.repo.Get(link)
	if err != nil {
		return nil, fmt.Errorf("/repository: %w", err)
	}
	return short, nil
}

func (serv *linkServise) Create(links *entity.LinkCreate) error {
	var err error
	links.ShortLink = serv.utils.GenerateShortUrl()
	if err != nil {
		return fmt.Errorf("/utils: %w", err)
	}
	err = serv.repo.Create(links)
	if err != nil {
		return fmt.Errorf("/repository: %w", err)
	}
	return nil
}
