package utils

import (
	"go.uber.org/zap"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	SugarLogger = logger.Sugar()
}
