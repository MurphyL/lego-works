package account

import (
	"github.com/MurphyL/lego-works/pkg/dal"
)

type Account struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	PersonID uint   `gorm:"uniqueIndex" json:"personId"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func (a Account) TableName() string {
	return "sys_account"
}

func GetAccount(id string) (any, error) {
	acc := &Account{}
	dao := dal.GetDefaultRepo()
	return acc, dal.RetrieveOne(dao, acc, id)
}

func GetAccountHashPassword(username string) (string, error) {
	acc := &Account{}
	dao := dal.GetDefaultRepo()
	if err := dal.RetrieveOne(dao, acc, "username = ?", username); err == nil {
		return acc.Password, nil
	} else {
		return "", err
	}
}
