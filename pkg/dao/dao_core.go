package dao

import (
	"log"

	"github.com/MurphyL/lego-works/pkg/dao/internal/domain"
	"gorm.io/gorm"
)

type Repo interface {
	ApplyRetrieveOne(dest any, h RetrieveOne[any]) error
	ApplyRetrieveAll(dest any, h RetrieveAll[any]) error
	ApplyRetrieveWithPaging(dest any, h RetrieveAll[any]) error
}

type Model struct{}

type RetrieveOne[T any] func(T) (string, []any)
type RetrieveAll[T any] func(T) (string, []any)

func NewRepo(dial gorm.Dialector) domain.Repo {
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Println("连接数据库出错：", err.Error())
	}
	return &domain.GormRepo{DB: db}
}
