package entity

import "time"

type Channel struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Ordering    int       `json:"ordering"`
	Image       string    `json:"image"`
	Logo        string    `json:"logo"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   int64     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ChannelResponse) TableName() string {
	return "channels"
}

type ChannelResponse struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
