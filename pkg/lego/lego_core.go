package lego

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/MurphyL/lego-works/pkg/lego/internal/logger"
	"github.com/MurphyL/lego-works/pkg/lego/internal/rest"
)

var zapLogger *zap.Logger

func init() {
	zapLogger = logger.NewZapLogger()
}

func LookupDefaultDatasourceName() (string, bool) {
	return os.LookupEnv("GO_DSN_MYSQL")
}

func NewDomainRef(subdomain ...string) string {
	return strings.Join(subdomain, "#")
}

func NewSugarSugar() *zap.SugaredLogger {
	return zapLogger.Sugar()
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
