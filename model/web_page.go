package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebPage struct {
	ID        uint            `gorm:"column:id"`
	Code      string          `gorm:"column:code"`
	Name      string          `gorm:"column:name"`
	IsSystem  uint8           `gorm:"column:is_system"`
	Content   string          `gorm:"column:content"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`

	SEO *WebSEO `gorm:"foreignKey:ChannelID;references:ID"`
}

const TableWebPage = "web_page"

func (WebPage) TableName() string {
	return TableWebPage
}

const (
	WebPageOfIsSystemYes = 1
	WebPageOfIsSystemNo  = 2
)
