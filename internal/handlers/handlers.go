package handlers

import (
	"short_link_servise/internal/servise"
)

func NewHandlers(serv *servise.Servises) *Handlers {
	return &Handlers{
		Links: NewLinksHandler(serv.Links),
	}
}
