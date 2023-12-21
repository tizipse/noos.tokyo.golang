package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebHTML struct {
	ID        uint            `gorm:"column:id"`
	Channel   string          `gorm:"column:channel"`
	ChannelID string          `gorm:"column:channel_id"`
	Content   string          `gorm:"column:content"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableWebHTML = "web_html"

func (WebHTML) TableName() string {
	return TableWebHTML
}
