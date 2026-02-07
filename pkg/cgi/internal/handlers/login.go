package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MurphyL/lego-works/pkg/iam"
	"github.com/MurphyL/lego-works/pkg/lego"
)

var logger = lego.NewSugarSugar()

type GetHashPassword func(string) (string, error)

type LoginArgs struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CaptchaCode string `json:"captchaCode"`
	CaptchaKey  string `json:"captchaKey"`
}

func NewLoginHandler(getHashPassword GetHashPassword) func(*gin.Context) {
	return func(c *gin.Context) {
		args := &LoginArgs{}
		if c.BindJSON(args); !args.ValidRequest() {
			logger.Warn("登录出错，请求参数非法")
			c.JSON(http.StatusBadRequest, lego.NewResultViaMessage("非法的登录请求"))
		} else if hash, err := getHashPassword(args.Username); err != nil {
			logger.Warnf("登录出错，用户（username=%s）不存在", args.Username)
			c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage("用户不存在"))
		} else if iam.CompareHashPassword(hash, args) {
			logger.Infof("用户（username=%s）登录成功", args.Username)
			c.JSON(http.StatusOK, lego.NewResultViaPayload(args.Username))
		} else {
			logger.Warnf("登录出错，用户（username=%s）密码不正确", args.Username)
			c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage("用户密码错误"))
		}
	}
}

func LogoutHandler(c *gin.Context) {

}

func RefreshTokenHandler(c *gin.Context) {

}

func (a *LoginArgs) ValidRequest() bool {
	return true
}

func (a *LoginArgs) GetLoginUsername() string {
	return a.Username
}

func (a *LoginArgs) GetLoginPassword() string {
	return a.Password
}
