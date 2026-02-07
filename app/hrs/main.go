package main

import (
	"strings"

	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/MurphyL/lego-works/app/hrs/biz/account"
	"github.com/MurphyL/lego-works/pkg/cgi"
	"github.com/MurphyL/lego-works/pkg/dal"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var logger = lego.NewSugarSugar()

func main() {

	ctx := context.Background()
	dal.InitDefaultRepo(func(dsn string) gorm.Dialector {
		_, host, _ := strings.Cut(dsn, "@")
		logger.Info("尝试连接主数据库：", host)
		return mysql.Open(dsn)
	})
	app := cgi.NewRestApp(ctx)
	app.UseAuthHandlers("/api/v1/auth", account.GetAccountHashPassword)
	app.RetrieveOne("/api/v1/accounts/:id", account.GetAccount)
	app.Serve(":3000")
}
