package main

import (
	"time"

	"github.com/jungle85gopy/rotlogs/sizerot"
)

func main() {
	sizerot.BaseFileName = "log/access.log"
	sizerot.EnableDynamic = true
	// sizerot.JSONFormat = false
	sizerot.LogLevel = sizerot.DebugLevel
	logger := sizerot.Rotate()

	for {
		logger.Debug("it is a debug info!")
		logger.Info("it is info level log")
		logger.Warn("it is warning level log")
		logger.Error("it is warning level log")
		time.Sleep(500 * time.Millisecond)
	}
}
