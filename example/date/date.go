package main

import (
	"log"

	rotlog "github.com/jungle85gopy/rotlogs/daterot"
)

var (
	logger = rotlog.LoggerPtr.Logger
)

func main() {
	rotlog.BaseFileName = "logs/access.log"
	rotlog.BaseLinkName = rotlog.BaseFileName
	rotlog.RotateHour = 1
	// rotlog.JSONFormat = false
	rotlog.LogLevel = rotlog.DebugLevel
	var err error
	logger, err = rotlog.Rotate()
	if err != nil {
		log.Fatalf("create rotate file error:%s\n", err)
	}

	// defined in test.go for test of use logger
	testRot()
}
