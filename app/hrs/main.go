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
	dsn := os.Getenv("GO_DSN_MYSQL")
	dal.InitDefaultRepo(dsn, mysql.Open)
	idp := iam.NewIdentityProvider(ctx)
	app := cgi.NewRestApp(ctx)
	app.UseAuthHandlers("/api/v1/auth", idp)
	app.RetrieveOne("/api/v1/accounts/:id", account.GetUserInfo)
	app.Serve(":3000")
}
