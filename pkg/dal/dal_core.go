package dal

import (
	"os"

	"gorm.io/gorm"

	"github.com/MurphyL/lego-works/pkg/lego"
)

var (
	logger = lego.NewSugarSugar()
	// 默认持久化数据源
	defaultRepo *gorm.DB
)

type RepoOption func(*gorm.Config)

func NewRepo(dial gorm.Dialector, withOpts ...RepoOption) *gorm.DB {
	config := &gorm.Config{}
	for _, withOpt := range withOpts {
		withOpt(config)
	}
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		logger.Panicln("连接数据库出错：", err.Error())
	}
	return db
}

func InitDefaultRepo(makeDial func(dsn string) gorm.Dialector, withOpts ...RepoOption) {
	dsn := os.Getenv("GO_DSN_MYSQL")
	dial := makeDial(dsn)
	defaultRepo = NewRepo(dial, withOpts...)
}

func GetDefaultRepo() *gorm.DB {
	return defaultRepo
}

func GetRepo(key string) *gorm.DB {
	return defaultRepo
}

func CreateOne(db *gorm.DB, dest interface{}) error {
	return db.Create(dest).Error
}

func UpdateOne(db *gorm.DB, dest interface{}, args ...interface{}) error {
	return db.Update("", dest).Error
}

func RetrieveOne(db *gorm.DB, dest interface{}, args ...interface{}) error {
	return db.Take(dest, args...).Error
}

func RetrieveAll(db *gorm.DB, dest interface{}, args ...interface{}) error {
	return db.Find(dest, args...).Error
}

func RetrieveWithPaging(db *gorm.DB, dest interface{}, args ...interface{}) error {
	return db.Find(dest, args...).Error
}

func CreateOrUpdate(db *gorm.DB, dest interface{}) error {
	return db.Create(dest).Error
}
