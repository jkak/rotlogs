package main

import (
	"time"
)

func testRot() {
	for {
		logger.Debug("it is a debug info!")
		logger.Info("it is info level log")
		logger.Warn("it is warning level log")
		logger.Error("it is warning level log")
		time.Sleep(500 * time.Millisecond)
	}
}
