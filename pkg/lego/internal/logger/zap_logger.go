package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	// 编码配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 文件拆分
	lumberJackSyncer := &lumberjack.Logger{
		Filename:   "./hrs.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackSyncer)
	// 默认等级
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	// 日志门面
	return zap.New(core, zap.AddCaller())
}
