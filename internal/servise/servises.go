package servise

import (
	"short_link_servise/internal/repository"
)

func NewServises(repos *repository.Repos) *Servises {
	return &Servises{
		Links: NewLinksServise(repos.Links),
	}
}
