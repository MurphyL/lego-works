package dal

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/MurphyL/lego-works/pkg/dal/internal/rdbms"
	"github.com/MurphyL/lego-works/pkg/lego"
)

const (
	defaultDsn = "default"
)

var (
	logger     = lego.NewSugarSugar()
	dalContext = context.Background()
)

type Repo = *rdbms.GormRepo

type RepoOption func(*gorm.Config)

func GetDefaultRepo() Repo {
	return GetRepo(defaultDsn)
}

func GetRepo(key string) Repo {
	repoKey := makeRepoKey(key)
	repo := dalContext.Value(repoKey)
	if nil == repo {
		logger.Panicf("未找到数据仓库（%s）", repoKey)
		return nil
	}
	return repo.(Repo)
}

func SetDefaultRepo(repo Repo) {
	dalContext = context.WithValue(dalContext, makeRepoKey(defaultDsn), repo)
}

// NewGormRepo Gorm 数据仓库
func NewGormRepo(dsn string, openDatabase func(dsn string) gorm.Dialector, withOpts ...RepoOption) Repo {
	_, host, _ := strings.Cut(dsn, "@")
	logger.Infof("尝试连接数据库（%s）……", host)
	config := &gorm.Config{}
	for _, withOpt := range withOpts {
		withOpt(config)
	}
	db, err := gorm.Open(openDatabase(dsn), &gorm.Config{})
	if err != nil {
		logger.Panicln("连接数据库出错：", host, err.Error())
	}
	key := makeRepoKey(uuid.NewString())
	repo := &rdbms.GormRepo{DB: db, Key: key}
	dalContext = context.WithValue(dalContext, key, repo)
	return repo
}

func makeRepoKey(key string) string {
	return lego.NewDomainRef("repo", key)
}
