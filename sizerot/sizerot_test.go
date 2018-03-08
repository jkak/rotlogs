package sizerot

import (
	"testing"
	"time"
)

func TestRotate(t *testing.T) {
	BaseFileName = "log/access.log"
	// JSONFormat = false
	EnableDynamic = true

	LogLevel = DebugLevel
	logger := Rotate()

	for {
		logger.Debug("it is debug level log")
		logger.Info("it is info level log")
		logger.Warn("it is warning level log")
		logger.Error("it is warning level log")
		// logger.Fatal("it is fatal level")
		time.Sleep(500 * time.Millisecond)
	}
}
