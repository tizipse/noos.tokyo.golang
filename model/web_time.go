package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebTime struct {
	ID        uint            `gorm:"column:id"`
	Name      string          `gorm:"column:name"`
	Content   string          `gorm:"column:content"`
	Status    string          `gorm:"column:status"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableWebTime = "web_time"

func (WebTime) TableName() string {
	return TableWebTime
}

const (
	WebTimeOfStatusByOpen  = "open"
	WebTimeOfStatusByClose = "close"
)
