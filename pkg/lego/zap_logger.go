package lego

import (
	"go.uber.org/zap"

	"github.com/MurphyL/lego-works/pkg/lego/internal/logger"
)

var zapLogger *zap.Logger

func init() {
	zapLogger = logger.NewZapLogger()
}

func NewSugarSugar() *zap.SugaredLogger {
	return zapLogger.Sugar()
}
