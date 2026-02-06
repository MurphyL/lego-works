package main

import (
	"log"
	"testing"

	"github.com/MurphyL/lego-works/pkg/dao"
	"gorm.io/driver/mysql"
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
	return "sys_user_account"
}

func TestLucky(t *testing.T) {
	conn := mysql.Open("yzm_hyqf_rw:VOWB@ypTw4ayzfqx2Vq@etmut2JFmNtg@tcp(192.168.33.233:3306)/test_yzm_hyqf?charset=utf8mb4&parseTime=True&loc=Local")
	repo := dao.NewRepo(conn)
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
