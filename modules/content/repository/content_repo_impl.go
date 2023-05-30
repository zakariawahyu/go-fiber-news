package repository

import (
	"github.com/zakariawahyu/go-fiber-news/entity"
	"github.com/zakariawahyu/go-fiber-news/helpers"
	"gorm.io/gorm"
)

type ContentRepositoryImpl struct {
	Conn *gorm.DB
}

func NewContentRepository(Conn *gorm.DB) ContentRepository {
	return &ContentRepositoryImpl{Conn}
}

func (repo *ContentRepositoryImpl) GetBySlug(slug string) (entity.Content, error) {
	var content entity.Content

	err := repo.Conn.Where("slug = ?", slug).Preload("User").Preload("Region").Preload("Channel").Preload("SubChannel").Preload("Tags").Preload("Topics").Preload("Reporters").Preload("SubPhotos").First(&content).Error
	if err != nil {
		return content, helpers.ErrNotFound
	}

	return content, nil
}
