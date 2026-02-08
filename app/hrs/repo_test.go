package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"

	"github.com/MurphyL/lego-works/pkg/dal"
)

type Account struct {
	Id       uint64
	Username string
	Password string
	RealName string
	UserCode string
	PhoneNum string
}

func (a *Account) TableName() string {
	if hostname, _ := os.Hostname(); "lucky" == hostname {
		return "sys_user_account"
	}
	return "sys_account"
}

func TestRepo(t *testing.T) {
	dsn := os.Getenv("GO_DSN_MYSQL")
	conn := mysql.Open(dsn)
	repo := dal.NewGormRepo(conn)
	acc := &Account{}
	if err := repo.RetrieveOne(acc, 1); err == nil {
		log.Println("用户查询完成：", acc)
	} else {
		log.Println("用户查询出错：", err.Error())
	}
	records := make([]Account, 0)
	if err := repo.RetrieveAll(&records, []any{1, 2, 3, 4, 5}); err == nil {
		log.Println("用户查询完成：", records)
	} else {
		log.Println("用户查询出错：", err.Error())
	}
	records = make([]Account, 0)
	if err := repo.RetrieveAll(&records, []any{1, 2, 3, 4, 5}); err == nil {
		log.Println("用户查询完成：", records)
	} else {
		log.Println("用户查询出错：", err.Error())
	}
}
