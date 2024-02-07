package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebOriginal struct {
	ID        string          `gorm:"column:id"`
	Name      string          `gorm:"column:name"`
	Thumb     string          `gorm:"column:thumb"`
	INS       string          `gorm:"column:ins"`
	Summary   string          `gorm:"column:summary"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`

	SEO  *WebSEO  `gorm:"foreignKey:ChannelID;references:ID"`
	HTML *WebHTML `gorm:"foreignKey:ChannelID;references:ID"`
}

const TableWebOriginal = "web_original"

func (WebOriginal) TableName() string {
	return TableWebOriginal
}
