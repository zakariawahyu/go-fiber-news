package services

import (
	"github.com/zakariawahyu/go-fiber-news/entity"
	"github.com/zakariawahyu/go-fiber-news/exception"
	"github.com/zakariawahyu/go-fiber-news/modules/content/repository"
)

type ContentServicesImpl struct {
	contentRepo repository.ContentRepository
}

func NewContentServices(repo repository.ContentRepository) ContentServices {
	return &ContentServicesImpl{
		contentRepo: repo,
	}
}

func (serv *ContentServicesImpl) GetBySlug(slug string) entity.ContentResponse {
	res, err := serv.contentRepo.GetBySlug(slug)
	exception.PanicIfNeeded(err)

	return entity.NewContentResponse(res)
}
