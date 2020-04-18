package logutil

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLogger(t *testing.T) {
	log := GetLogger(zapcore.DebugLevel, "app.log")
	log.Debug("Init logger successed.")
	log.Info("Print log info with level 'INFO'.")
}
