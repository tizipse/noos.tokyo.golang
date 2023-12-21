package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebSEO struct {
	ID          uint            `gorm:"column:id"`
	Channel     string          `gorm:"column:channel"`
	ChannelID   string          `gorm:"column:channel_id"`
	Title       string          `gorm:"column:title"`
	Keyword     string          `gorm:"column:keyword"`
	Description string          `gorm:"column:description"`
	CreatedAt   carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableWebSEO = "web_seo"

func (WebSEO) TableName() string {
	return TableWebSEO
}
