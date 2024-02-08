package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebLink = "web_link"

type WebLink struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Name      string          `gorm:"column:name"`
	Summary   string          `gorm:"column:summary"`
	URL       string          `gorm:"column:url"`
	Order     uint8           `gorm:"column:order"`
	IsSystem  uint8           `gorm:"column:is_system"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (WebLink) TableName() string {
	return TableWebLink
}

const (
	WebLinkOfIsSystemYES = 1
	WebLinkOfIsSystemNO  = 2
)
