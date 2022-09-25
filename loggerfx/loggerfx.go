package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func LoggerProvider() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	return slogger
}

var Module = fx.Options(fx.Provide(LoggerProvider))
