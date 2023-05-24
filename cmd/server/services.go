package server

import (
	_contentServices "github.com/zakariawahyu/go-fiber-news/modules/content/services"
)

type Services struct {
	contentServices _contentServices.ContentServices
}

func NewServices(repo *Repository) *Services {
	return &Services{
		contentServices: _contentServices.NewContentServices(repo.ContentRepo),
	}
}
