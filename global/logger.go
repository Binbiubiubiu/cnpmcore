package global

import (
	"go.uber.org/zap"
)

var ZLog *zap.SugaredLogger

func init() {
	zap.NewDevelopment()
	log, _ := zap.NewDevelopment()
	ZLog = log.Sugar()
}
