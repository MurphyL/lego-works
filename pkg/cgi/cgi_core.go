package cgi

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MurphyL/lego-works/pkg/cgi/internal/domain"
	"github.com/MurphyL/lego-works/pkg/dal"
)

func NewRestApp(ctx context.Context, dao dal.Repo) *RestApp {
	gin.SetMode(gin.ReleaseMode)
	return &RestApp{
		ctx:    ctx,
		dao:    dao,
		engine: gin.New(),
	}
}

type RestApp struct {
	ctx    context.Context
	dao    dal.Repo
	engine *gin.Engine
}

func (a *RestApp) HandleRequest(endpoint string, handler func(dal.Repo, []byte) (any, error)) {
	a.engine.POST(endpoint, func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		if ret, err := handler(a.dao, body); err == nil {
			c.JSON(http.StatusOK, &domain.Result{Payload: ret})
		} else {
			c.JSON(http.StatusInternalServerError, domain.Result{Message: err.Error()})
		}
	})
}

func (a *RestApp) RetrieveOne(endpoint string, retriever func(dal.Repo, string) (any, error)) {
	a.engine.GET(endpoint, func(c *gin.Context) {
		refValue := c.Param(c.DefaultQuery("ref", "id"))
		if dest, err := retriever(a.dao, refValue); err == nil {
			c.JSON(http.StatusOK, dest)
		} else {
			c.JSON(http.StatusInternalServerError, domain.Result{Message: err.Error()})
		}
	})
}

func (a *RestApp) Serve(addr string) {
	a.engine.Use(gin.Recovery())
	srv := &http.Server{
		Addr:    addr,
		Handler: a.engine.Handler(),
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
