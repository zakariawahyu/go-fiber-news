package server

import (
	_contentRepository "github.com/zakariawahyu/go-fiber-news/modules/content/repository"
	"github.com/zakariawahyu/go-fiber-news/utils"
)

type Repository struct {
	ContentRepo _contentRepository.ContentRepository
}

func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		ContentRepo: _contentRepository.NewContentRepository(conn.Mysql),
	}
}
