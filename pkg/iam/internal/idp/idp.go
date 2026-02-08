package idp

import (
	"github.com/MurphyL/lego-works/pkg/dal"
)

type AccountManagerOption func(*AccountRepo)

func NewAccountRepo(withOpts ...AccountManagerOption) *AccountRepo {
	idp := &AccountRepo{}
	for _, withOpt := range withOpts {
		withOpt(idp)
	}
	if nil == idp.Repo {
		idp.Repo = dal.GetDefaultRepo()
	}
	return idp
}

type AccountRepo struct {
	dal.Repo
}

func (idp *AccountRepo) LoadAccountInfo(dest any, username string) error {
	return dal.GetDefaultRepo().Take(dest, "username = ?", username).Error
}
