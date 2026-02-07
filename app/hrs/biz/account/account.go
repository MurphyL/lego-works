package account

import (
	"github.com/MurphyL/lego-works/pkg/dal"
	"github.com/MurphyL/lego-works/pkg/iam"
)

func GetUserInfo(id string) (any, error) {
	acc := iam.NewAccount()
	return acc, dal.GetDefaultRepo().Take(acc, id).Error
}
