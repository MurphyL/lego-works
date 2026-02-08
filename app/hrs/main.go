package main

import (
	"os"

	"golang.org/x/net/context"
	"gorm.io/driver/mysql"

	"github.com/MurphyL/lego-works/app/hrs/biz/account"
	"github.com/MurphyL/lego-works/pkg/cgi"
	"github.com/MurphyL/lego-works/pkg/dal"
	"github.com/MurphyL/lego-works/pkg/iam"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var logger = lego.NewSugarSugar()

func main() {
	ctx := context.Background()
	dsn, dsnDefined := lego.LookupDefaultDatasourceName()
	if !dsnDefined {
		logger.Panicln("未找到数据源配置")
		os.Exit(1)
	}
	repo := dal.NewGormRepo(dsn, mysql.Open)
	dal.SetDefaultRepo(repo)
	idp := iam.NewIdentityProvider(ctx)
	app := cgi.NewRestApp(ctx)
	app.UseAuthHandlers("/api/v1/auth", idp)
	app.RetrieveOne("/api/v1/accounts/:id", account.GetUserInfo)
	app.Serve(":3000")
}
