package main

import (
	"log"
	"time"

	"github.com/jungle85gopy/rotlogs/daterot"
)

func main() {
	daterot.BaseFileName = "log/access.log"
	daterot.BaseLinkName = daterot.BaseFileName
	daterot.RotateHour = 1
	// daterot.JSONFormat = false
	daterot.LogLevel = daterot.DebugLevel
	logger, err := daterot.Rotate()
	if err != nil {
		log.Fatalf("create rotate file error:%s\n", err)
	}
	for {
		logger.Debug("it is a debug info!")
		logger.Info("it is info level log")
		logger.Warn("it is warning level log")
		logger.Error("it is warning level log")
		time.Sleep(500 * time.Millisecond)
	}
}
