package dal

import (
	"log"

	"gorm.io/gorm"

	"github.com/MurphyL/lego-works/pkg/dal/internal/domain"
)

type Repo interface {
	CreateOne(dest interface{}) error
	UpdateOne(dest interface{}, args ...interface{}) error
	RetrieveOne(dest interface{}, args ...interface{}) error
	RetrieveAll(dest interface{}, args ...interface{}) error
	RetrieveWithPaging(dest interface{}, args ...interface{}) error
	CreateOrUpdate(dest interface{}) error
}

func NewRepo(dial gorm.Dialector) Repo {
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Println("连接数据库出错：", err.Error())
	}
	return &domain.GormRepo{DB: db}
}
