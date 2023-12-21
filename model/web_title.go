package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebTitle struct {
	ID        uint            `gorm:"column:id"`
	Name      string          `gorm:"column:name"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableWebTitle = "web_title"

func (WebTitle) TableName() string {
	return TableWebTitle
}
