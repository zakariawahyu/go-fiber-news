package entity

import (
	"time"
)

type Content struct {
	ID               int64               `gorm:"primaryKey" json:"id"`
	Type             string              `json:"type"`
	TypeID           int64               `json:"type_id"`
	TypeChildID      int64               `json:"type_child_id"`
	SuplemenID       int64               `json:"suplemen_id"`
	UpperTitle       string              `gorm:"column:uppertitle" json:"upper_title"`
	Title            string              `json:"title"`
	Slug             string              `json:"slug"`
	Excerpt          string              `json:"excerpt"`
	Content          string              `json:"content"`
	Caption          string              `json:"caption"`
	Image            string              `json:"image"`
	Thumbnail        string              `json:"thumbnail"`
	HeadlineType     int64               `json:"headline_type"`
	IsActive         bool                `json:"is_active"`
	IsFeatured       bool                `json:"is_featured"`
	IsEditorChoice   bool                `json:"is_editor_choice"`
	IsNational       bool                `json:"is_national"`
	Token            string              `json:"token"`
	AdsPosition      int                 `json:"ads_position"`
	AdsExpiredDate   time.Time           `json:"ads_expired_date"`
	UrlVideo         string              `json:"url_video"`
	IsPopular        bool                `json:"is_popular"`
	IsAdult          bool                `json:"is_adult"`
	IsComment        bool                `json:"is_comment"`
	LocationCity     string              `json:"location_city"`
	LocationDistrict string              `gorm:"column:location_citydistrict" json:"location_district"`
	LocationSuburb   string              `json:"location_suburb"`
	CreatedBy        int64               `gorm:"not null" json:"created_by"`
	User             *UserResponse       `gorm:"foreignKey:CreatedBy" json:"user"`
	Region           *RegionResponse     `gorm:"foreignKey:TypeID" json:"region"`
	Channel          *ChannelResponse    `gorm:"foreignKey:TypeID" json:"channel"`
	SubChannel       *SubChannelResponse `gorm:"foreignKey:TypeChildID" json:"sub_channel"`
	Tag              []Tag               `gorm:"many2many:content_has_tags;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:TagID" json:"tags"`
	Topic            []Topic             `gorm:"many2many:content_has_topics;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:TopicID" json:"topics"`
	Reporter         []Reporter          `gorm:"many2many:content_has_reporters;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:ReporterID" json:"reporters"`
	SubPhoto         []SubPhoto          `gorm:"foreignKey:ContentID" json:"sub_photos"`
	PublishedDate    time.Time           `json:"published_date"`
	CreatedAt        time.Time           `gorm:"column:created" json:"created_at"`
	UpdatedAt        time.Time           `gorm:"column:modified" json:"updated_at"`
}

type ContentResponse struct {
	ID            int64               `gorm:"primaryKey" json:"id"`
	Type          string              `json:"type"`
	TypeID        int64               `json:"type_id"`
	TypeChildID   int64               `json:"type_child_id"`
	SuplemenID    int64               `json:"suplemen_id"`
	UpperTitle    string              `gorm:"column:uppertitle" json:"upper_title"`
	Title         string              `json:"title"`
	Slug          string              `json:"slug"`
	Excerpt       string              `json:"excerpt"`
	Content       string              `json:"content"`
	Caption       string              `json:"caption"`
	Image         string              `json:"image"`
	Thumbnail     string              `json:"thumbnail"`
	UrlVideo      string              `json:"url_video"`
	IsPopular     bool                `json:"is_popular"`
	IsAdult       bool                `json:"is_adult"`
	IsComment     bool                `json:"is_comment"`
	CreatedBy     int64               `gorm:"not null" json:"created_by"`
	User          *UserResponse       `gorm:"foreignKey:CreatedBy" json:"user"`
	Region        *RegionResponse     `gorm:"foreignKey:TypeID" json:"region"`
	Channel       *ChannelResponse    `gorm:"foreignKey:TypeID" json:"channel"`
	SubChannel    *SubChannelResponse `gorm:"foreignKey:TypeChildID" json:"sub_channel"`
	Tag           []Tag               `gorm:"many2many:content_has_tags;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:TagID" json:"tags"`
	Topic         []Topic             `gorm:"many2many:content_has_topics;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:TopicID" json:"topics"`
	Reporter      []Reporter          `gorm:"many2many:content_has_reporters;ForeignKey:ID;joinForeignKey:ContentID;References:ID;joinReferences:ReporterID" json:"reporters"`
	SubPhoto      []SubPhoto          `gorm:"foreignKey:ContentID" json:"sub_photos"`
	PublishedDate time.Time           `json:"published_date"`
	UpdatedAt     time.Time           `gorm:"column:modified" json:"updated_at"`
}

func (Content) TableName() string {
	return "contents_new"
}

func (ContentResponse) TableName() string {
	return "contents_new"
}

func NewContentResponse(content Content) ContentResponse {
	return ContentResponse{
		ID:            content.ID,
		Type:          content.Type,
		TypeChildID:   content.TypeChildID,
		SuplemenID:    content.SuplemenID,
		UpperTitle:    content.UpperTitle,
		Title:         content.Title,
		Slug:          content.Slug,
		Excerpt:       content.Excerpt,
		Content:       content.Content,
		Caption:       content.Caption,
		Image:         content.Image,
		Thumbnail:     content.Thumbnail,
		UrlVideo:      content.UrlVideo,
		IsPopular:     content.IsPopular,
		IsAdult:       content.IsAdult,
		IsComment:     content.IsComment,
		PublishedDate: content.PublishedDate,
		UpdatedAt:     content.UpdatedAt,
		User:          content.User,
		Region:        content.Region,
		Channel:       content.Channel,
		SubChannel:    content.SubChannel,
		Tag:           content.Tag,
		Topic:         content.Topic,
		Reporter:      content.Reporter,
		SubPhoto:      content.SubPhoto,
	}
}
