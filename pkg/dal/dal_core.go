package dal

import (
	"strings"

	"gorm.io/gorm"

	"github.com/MurphyL/lego-works/pkg/dal/internal/rdbms"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var (
	logger = lego.NewSugarSugar()
	// 默认持久化数据源
	defaultRepo *rdbms.GormRepo
)

type Repo = *rdbms.GormRepo

type RepoOption func(*gorm.Config)

func NewGorm(dial gorm.Dialector, withOpts ...RepoOption) Repo {
	config := &gorm.Config{}
	for _, withOpt := range withOpts {
		withOpt(config)
	}
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		logger.Panicln("连接数据库出错：", err.Error())
	}
	return &rdbms.GormRepo{DB: db}
}

func InitDefaultRepo(dsn string, openDatabase func(dsn string) gorm.Dialector, withOpts ...RepoOption) {
	_, host, _ := strings.Cut(dsn, "@")
	logger.Info("尝试连接主数据库：", host)
	dial := openDatabase(dsn)
	defaultRepo = NewGorm(dial, withOpts...)
}

func GetDefaultRepo() Repo {
	return defaultRepo
}
