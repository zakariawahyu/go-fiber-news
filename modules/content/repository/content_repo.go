package repository

import "github.com/zakariawahyu/go-fiber-news/entity"

type ContentRepository interface {
	GetBySlug(slug string) (entity.Content, error)
}
