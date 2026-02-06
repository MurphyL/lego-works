package domain

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App interface {
	Serve(string)
}

func NewApp() App {
	gin.SetMode(gin.ReleaseMode)
	return &internalRestApp{
		Engine: gin.New(),
	}
}

type internalRestApp struct {
	*gin.Engine
}

func (a internalRestApp) Serve(addr string) {
	a.Use(gin.Recovery())
	srv := &http.Server{
		Addr:    addr,
		Handler: a.Handler(),
	}
	// 启动服务器协程
	go func() {
		if err := srv.ListenAndServe(); err == nil {
			log.Fatalln("Server Shutdown:", err)
		}
	}()
	log.Println("Server started:", addr)
	// 监听中断信号并触发优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exited")
}
