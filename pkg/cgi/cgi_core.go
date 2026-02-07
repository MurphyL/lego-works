package cgi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MurphyL/lego-works/pkg/cgi/internal/handlers"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var logger = lego.NewSugarSugar()

func NewRestApp(ctx context.Context) *RestServer {
	gin.SetMode(gin.ReleaseMode)
	return &RestServer{
		ctx:    ctx,
		router: gin.New(),
	}
}

type RestServer struct {
	ctx    context.Context
	router *gin.Engine
}

func (a *RestServer) UseAuthHandlers(endpoint string, getHashPassword handlers.GetHashPassword) {
	logger.Info("正在注册授权模块……")
	r := a.router.Group(endpoint)
	r.POST("/login", handlers.NewLoginHandler(getHashPassword))
	r.GET("/logout", handlers.LogoutHandler)
	a.router.Use(func(c *gin.Context) {
		// 验证token
		// 验证权限
		c.Next()
	})
}

func (a *RestServer) RetrieveOne(endpoint string, retriever func(string) (any, error)) {
	a.router.GET(endpoint, func(c *gin.Context) {
		refValue := c.Param(c.DefaultQuery("ref", "id"))
		if dest, err := retriever(refValue); err == nil {
			c.JSON(http.StatusOK, lego.NewResultViaPayload(dest))
		} else {
			c.JSON(http.StatusInternalServerError, lego.NewResultViaError(err))
		}
	})
}

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
