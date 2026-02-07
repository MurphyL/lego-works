package idp

import (
	"github.com/MurphyL/lego-works/pkg/dal"
)

type AccountManagerOption func(*AccountManager)

func NewAccountRepo(withOpts ...AccountManagerOption) *AccountManager {
	idp := &AccountManager{}
	for _, withOpt := range withOpts {
		withOpt(idp)
	}
	if nil == idp.Repo {
		idp.Repo = dal.GetDefaultRepo()
	}
	return idp
}

type AccountManager struct {
	dal.Repo
}

func (idp *AccountManager) LoadAccountInfo(dest any, username string) error {
	return dal.GetDefaultRepo().Take(dest, "username = ?", username).Error
}
