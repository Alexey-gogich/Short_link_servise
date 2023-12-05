package handlers

import "short_link_servise/internal/servise"

type linksHandler struct {
	serv servise.LinksService
}

func NewLinksHandler(serv servise.LinksService) *linksHandler {
	return &linksHandler{
		serv: serv,
	}
}

func (hand *linksHandler) Get() {

}

func (hand *linksHandler) Create() {

}
