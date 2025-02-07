package pkg

import (
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	Logger = logger.Sugar()
}
