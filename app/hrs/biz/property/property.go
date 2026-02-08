package property

import (
	"time"

	"github.com/MurphyL/lego-works/pkg/dal"
)

/**
 * Property 房产
 */

// Property 房产基础信息
type Property struct {
	ID            uint `gorm:"primaryKey"`
	PropertyTitle string
	PropertyArea  string // 合同面积
	SharedArea    string // 公摊面积
	Address       string // 房产未知
	Status        uint8
	CreatedTime   time.Time
	UpdatedTime   time.Time
}

func (*Property) TableName() string {
	return "hrs_property"
}

func GetProperty(id string) (*Property, error) {
	property := &Property{}
	return property, dal.GetDefaultRepo().Take(property, id).Error
}
