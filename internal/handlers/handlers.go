package handlers

import "short_link_servise/internal/servise"

type LinksHandler interface {
	Get()
	Create()
}

type Handlers struct {
	links LinksHandler
}

func NewHandlers(serv *servise.Servises) *Handlers {
	return &Handlers{
		links: NewLinksHandler(serv.Links),
	}
}
