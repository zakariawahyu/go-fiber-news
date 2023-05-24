package services

import "github.com/zakariawahyu/go-fiber-news/entity"

type ContentServices interface {
	GetBySlug(slug string) entity.ContentResponse
}
