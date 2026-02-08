package account

import (
	"github.com/MurphyL/lego-works/pkg/dal"
	"github.com/MurphyL/lego-works/pkg/iam"
)

// PersonInfo 公民信息
type PersonInfo struct {
	Id         uint64 `json:"id"`
	RealName   string `json:"realName"`
	IdCardType string `json:"idCardType"` // 证件类型
	IdCardNo   string `json:"idCardNo"`
}

func (a PersonInfo) TableName() string {
	return "base_person"
}

func GetUserInfo(id string) (any, error) {
	acc := iam.NewAccount()
	return acc, dal.GetDefaultRepo().Take(acc, id).Error
}
