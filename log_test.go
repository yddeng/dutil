package dutil

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger("log", "testLog")

	logger.Infoln("infoln message", 1)
	logger.Infof("%s : %d", "infof message", 2)
	logger.Debugln("Debugln message", 1)
	logger.Debugf("%s : %d", "Debugf message", 2)
	logger.Errorln("Errorln message", 1)
	logger.Errorf("%s : %d", "Errorf message", 2)

}