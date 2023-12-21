package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebMember struct {
	ID         string          `gorm:"column:id"`
	TitleID    uint            `gorm:"column:title_id"`
	Name       string          `gorm:"column:name"`
	Nickname   string          `gorm:"column:nickname"`
	Thumb      string          `gorm:"column:thumb"`
	INS        string          `gorm:"column:ins"`
	Order      uint8           `gorm:"column:order"`
	IsDelegate uint8           `gorm:"column:is_delegate"`
	IsEnable   uint8           `gorm:"column:is_enable"`
	CreatedAt  carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at"`

	Title *WebTitle `gorm:"foreignKey:ID;references:TitleID"`

	SEO  *WebSEO  `gorm:"foreignKey:ChannelID;references:ID"`
	HTML *WebHTML `gorm:"foreignKey:ChannelID;references:ID"`
}

const TableWebMember = "web_member"

func (WebMember) TableName() string {
	return TableWebMember
}
