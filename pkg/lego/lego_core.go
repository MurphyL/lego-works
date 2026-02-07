package lego

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/MurphyL/lego-works/pkg/lego/internal/logger"
	"github.com/MurphyL/lego-works/pkg/lego/internal/rest"
)

var zapLogger *zap.Logger

type CaptchaArgs rest.CaptchaArgs

func init() {
	zapLogger = logger.NewZapLogger()
}

func NewSugarSugar() *zap.SugaredLogger {
	return zapLogger.Sugar()
}

func NewCaptchaArgs(key, code string) *rest.CaptchaArgs {
	return &rest.CaptchaArgs{CaptchaKey: key, CaptchaCode: code}
}

func NewResult(ok bool, payload any, message string) *rest.Result {
	return &rest.Result{Success: ok, Payload: payload, Message: message}
}

func NewSuccessResult(payload any) *rest.Result {
	return NewResult(true, payload, "OK")
}

func NewResultViaError(err error) *rest.Result {
	return NewResult(false, nil, fmt.Sprintf("未知错误：%s", err.Error()))
}

func NewResultViaMessage(ok bool, msg string) *rest.Result {
	return NewResult(ok, nil, msg)
}
