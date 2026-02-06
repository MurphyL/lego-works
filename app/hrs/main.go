package main

import (
	"os"

	"golang.org/x/net/context"
	"gorm.io/driver/mysql"

	"github.com/MurphyL/lego-works/app/core/account"
	"github.com/MurphyL/lego-works/pkg/cgi"
	"github.com/MurphyL/lego-works/pkg/dal"
)

func main() {
	dsn := os.Getenv("GO_DSN_MYSQL")
	conn := mysql.Open(dsn)
	dao := dal.NewRepo(conn)
	ctx := context.Background()
	app := cgi.NewRestApp(ctx, dao)
	app.HandleRequest("/api/v1/auth/login", account.Login)
	app.RetrieveOne("/api/v1/accounts/:id", account.GetAccount)
	app.Serve(":3000")
}
