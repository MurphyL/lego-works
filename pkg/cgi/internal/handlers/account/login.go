package account

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/MurphyL/lego-works/pkg/iam"
	"github.com/MurphyL/lego-works/pkg/lego"
)

const (
	LoginSecretKey = "my_secret_key"
	RequestUserKey = "requestUser"
)

var logger = lego.NewSugarSugar()

func NewLoginHandler(idp iam.IdentityProvider) func(*gin.Context) {
	logger.Info("正在注册用户登录处理器……")
	return func(c *gin.Context) {
		var loginArgs, existsAcc = iam.NewLoginArgs(iam.PasswordActionTypeLogin), iam.NewAccount()
		if c.BindJSON(loginArgs); !loginArgs.ValidRequest() {
			logger.Warn("登录出错，请求参数非法")
			c.JSON(http.StatusBadRequest, lego.NewResultViaMessage(false, "非法的登录请求"))
		} else if err := idp.LoadAccountInfo(existsAcc, loginArgs.Username); err != nil {
			logger.Warnf("登录出错，用户（username=%s）不存在", loginArgs.Username)
			c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage(false, "用户不存在"))
		} else if captcha.VerifyString(loginArgs.CaptchaKey, loginArgs.CaptchaCode) {
			if loginArgs.CompareHashPassword(existsAcc.Password) {
				token := getAccountToken(uint64(existsAcc.ID), existsAcc.Username, "login")
				c.Header("X-Token", token)
				c.JSON(http.StatusOK, lego.NewSuccessResult(token))
			} else {
				logger.Warnf("登录出错，用户（username=%s）密码不正确", loginArgs.Username)
				c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage(false, "用户密码错误"))
			}
		} else {
			logger.Warnf("登录出错，验证码（captchaCode=%s）输入不正确", loginArgs.CaptchaCode)
			c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage(false, "验证码错误"))
		}
	}
}

func NewResetPasswordHandler(idp iam.IdentityProvider) func(*gin.Context) {
	return func(c *gin.Context) {
		var resetArgs, existsAcc = iam.NewLoginArgs(iam.PasswordActionTypeLogin), iam.NewAccount()
		if c.BindJSON(resetArgs); !resetArgs.ValidRequest() {
			logger.Warn("登录出错，请求参数非法")
			c.JSON(http.StatusBadRequest, lego.NewResultViaMessage(false, "非法的登录请求"))
		} else if err := idp.LoadAccountInfo(existsAcc, resetArgs.Username); err != nil {
			logger.Warnf("登录出错，用户（username=%s）不存在", resetArgs.Username)
			c.JSON(http.StatusInternalServerError, lego.NewResultViaMessage(false, "用户不存在"))
		} else {
			c.JSON(http.StatusOK, lego.NewResultViaMessage(true, "密码重置成功"))
		}
	}
}

func LogoutHandler(c *gin.Context) {

}

func AuthorizationHandler(c *gin.Context) {
	if authorization := c.Request.Header.Get("Authorization"); authorization != "" {
		if token, ok := strings.CutPrefix(authorization, "Bearer "); ok {
			claims := &jwt.RegisteredClaims{}
			_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(LoginSecretKey), nil
			})
			if err == nil {
				c.Set(RequestUserKey, claims.ID)
				c.Next()
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, lego.NewResultViaMessage(false, "请求未授权"))

}

func getAccountToken(uid uint64, username, subject string) string {
	// 定义声明
	claims := jwt.RegisteredClaims{
		ID:        strconv.FormatUint(uid, 10),
		Issuer:    username,
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名
	secretKey := []byte(LoginSecretKey)
	tokenString, _ := token.SignedString(secretKey)
	logger.Infof("用户（username=%s）登录成功，生成Token：%s", username, tokenString)
	return tokenString
}
