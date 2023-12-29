package model

import (
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type WebMenu struct {
	ID        uint            `gorm:"column:id"`
	Name      string          `gorm:"column:name"`
	Price     string          `gorm:"column:price"`
	Type      string          `gorm:"column:type"`
	Order     uint8           `gorm:"column:order"`
	IsEnable  uint8           `gorm:"column:is_enable"`
	CreatedAt carbon.DateTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`
}

const TableWebMenu = "web_menu"

func (WebMenu) TableName() string {
	return TableWebMenu
}

func (w WebMenu) ToType() string {

	val := ""

	switch w.Type {
	case WebMenuOfTypeKCut:
		val = WebMenuOfTypeVCut
	case WebMenuOfTypeKStyling:
		val = WebMenuOfTypeVStyling
	case WebMenuOfTypeKSpa:
		val = WebMenuOfTypeVSpa
	case WebMenuOfTypeKTreatment:
		val = WebMenuOfTypeVTreatment
	case WebMenuOfTypeKColor:
		val = WebMenuOfTypeVColor
	case WebMenuOfTypeKPerm:
		val = WebMenuOfTypeVPerm
	case WebMenuOfTypeKStraightPerm:
		val = WebMenuOfTypeVStraightPerm
	default:
		val = "Unknown"
	}

	return val
}

const (
	WebMenuOfTypeKCut          = "cut"
	WebMenuOfTypeKStyling      = "styling"
	WebMenuOfTypeKSpa          = "spa"
	WebMenuOfTypeKTreatment    = "treatment"
	WebMenuOfTypeKColor        = "color"
	WebMenuOfTypeKPerm         = "perm"
	WebMenuOfTypeKStraightPerm = "straight_perm"
)

const (
	WebMenuOfTypeVCut          = "CUT"
	WebMenuOfTypeVStyling      = "STYLING"
	WebMenuOfTypeVSpa          = "SPA"
	WebMenuOfTypeVTreatment    = "TREATMENT"
	WebMenuOfTypeVColor        = "COLOR"
	WebMenuOfTypeVPerm         = "PERM"
	WebMenuOfTypeVStraightPerm = "STRAIGHT PERM"
)
