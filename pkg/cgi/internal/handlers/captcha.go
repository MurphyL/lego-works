package handlers

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"

	"github.com/MurphyL/lego-works/pkg/lego"
)

const (
	// CollectNum The number of captchas created that triggers garbage collection used by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)
const (
	StdHeight = 80
	StdWidth  = 240
)

func init() {
	captcha.SetCustomStore(captcha.NewMemoryStore(CollectNum, Expiration))
}

func CaptchaHandler(c *gin.Context) {
	action := c.Query("action")
	if action == "" {
		c.JSON(http.StatusBadRequest, lego.NewResultViaMessage(false, "未指定请求类型"))
	} else {
		var content bytes.Buffer
		captchaId := captcha.New()
		if err := captcha.WriteImage(&content, captchaId, StdWidth, StdHeight); err == nil {
			c.Writer.Header().Set("Content-Type", "image/png")
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Writer.Header().Set("X-Captcha", captchaId)
			c.Writer.Write(content.Bytes())
		} else {
			c.JSON(http.StatusOK, lego.NewResultViaError(err))
		}
	}
}

func verifyCaptcha() {

}
