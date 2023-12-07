package inmemory

import (
	"short_link_servise/internal/repository"
)

func NewRepos(db Cache) *repository.Repos {
	return &repository.Repos{
		Links: NewLinksRepo(db),
	}
}
