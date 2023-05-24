package entity

type SubChannel struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type SubChannelResponse struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (SubChannelResponse) TableName() string {
	return "sub_channels"
}
