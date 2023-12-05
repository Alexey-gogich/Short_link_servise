package servise

import (
	"short_link_servise/internal/entity"
	"short_link_servise/internal/repository"
)

type LinksService interface {
	Get(link *entity.ShortLink) (*entity.Link, error)
	Create(links *entity.LinkCreate) error
}

type Servises struct {
	Links LinksService
}

func NewRepos(repos *repository.Repos) *Servises {
	return &Servises{
		Links: NewLinksServise(repos.Links),
	}
}
