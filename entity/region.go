package entity

type Region struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type RegionResponse struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (RegionResponse) TableName() string {
	return "regions"
}
