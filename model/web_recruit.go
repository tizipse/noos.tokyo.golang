package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

const TableWebRecruit = "web_recruit"

type WebRecruit struct {
	ID        uint            `gorm:"column:id;primaryKey"`
	Name      string          `gorm:"column:name"`
	Summary   string          `gorm:"column:summary"`
	URL       string          `gorm:"column:url"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

func (WebRecruit) TableName() string {
	return TableWebRecruit
}
