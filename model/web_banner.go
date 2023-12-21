package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebBanner = "web_banner"

type WebBanner struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Name      string          `gorm:"column:name"`
	Picture   string          `gorm:"column:picture"`
	Client    string          `gorm:"column:client"`
	Target    string          `gorm:"column:target"`
	URL       string          `gorm:"column:url"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (WebBanner) TableName() string {
	return TableWebBanner
}
