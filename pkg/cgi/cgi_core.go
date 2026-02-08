package cgi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/MurphyL/lego-works/pkg/cgi/internal/handlers"
	"github.com/MurphyL/lego-works/pkg/cgi/internal/handlers/account"
	"github.com/MurphyL/lego-works/pkg/iam"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var logger = lego.NewSugarSugar()

func NewRestApp(ctx context.Context) *RestServer {
	gin.SetMode(gin.ReleaseMode)
	// 创建路由管理器
	router := gin.New()
	// CORS 设置
	router.Use(cors.Default())
	// 使用 Logger 中间件
	router.Use(gin.Logger())
	// 使用 Recovery 中间件
	router.Use(gin.Recovery())
	// 创建服务器
	return &RestServer{
		ctx:    ctx,
		router: router,
	}
}

type RestServer struct {
	ctx    context.Context
	router *gin.Engine
}

// UseAuthHandlers 添加授权模块
func (a *RestServer) UseAuthHandlers(endpoint string, idp iam.IdentityProvider) {
	logger.Info("正在注册授权模块……")
	r := a.router.Group(endpoint)
	r.POST("/login", account.NewLoginHandler(idp))
	r.POST("/reset-password", account.NewResetPasswordHandler(idp))
	r.GET("/logout", account.LogoutHandler)
	r.GET("/captcha", handlers.CaptchaHandler)
	a.router.Use(func(c *gin.Context) {
		// 验证token
		// 验证权限
		c.Next()
	})
}

// RetrieveOne 获取单个对象
func (a *RestServer) RetrieveOne(endpoint string, retriever func(string) (any, error)) {
	a.router.GET(endpoint, account.AuthorizationHandler, func(c *gin.Context) {
		refValue := c.Param(c.DefaultQuery("ref", "id"))
		if dest, err := retriever(refValue); err == nil {
			c.JSON(http.StatusOK, lego.NewSuccessResult(dest))
		} else {
			c.JSON(http.StatusInternalServerError, lego.NewResultViaError(err))
		}
	})
}

// UpdateOne 修改单个对象
func (a *RestServer) UpdateOne(endpoint string, handler func(string) (any, error)) {
	a.router.PUT(endpoint, account.AuthorizationHandler, func(c *gin.Context) {
		refValue := c.Param(c.DefaultQuery("ref", "id"))
		if dest, err := handler(refValue); err == nil {
			c.JSON(http.StatusOK, lego.NewSuccessResult(dest))
		} else {
			c.JSON(http.StatusInternalServerError, lego.NewResultViaError(err))
		}
	})
}

// Serve 启动携程并运行服务器
func (a *RestServer) Serve(addr string) {
	a.router.Use(gin.Recovery())
	srv := &http.Server{
		Addr:    addr,
		Handler: a.router.Handler(),
	}

	// 启动服务器协程
	go func() {
		if err := srv.ListenAndServe(); err == nil {
			logger.Info("Server Shutdown:", err)
		}
	}()
	logger.Info("Server started:", addr)

	// 监听中断信号并触发优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Info("Server Shutdown:", err)
	}
	logger.Info("Server exited")
}
