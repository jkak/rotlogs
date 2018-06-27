package main

import (
	"log"

	rotlog "github.com/jkak/rotlogs/daterot"
)

var (
	logger = rotlog.LoggerPtr.Logger
)

func main() {
	rotlog.BaseFileName = "logs/access.log"
	rotlog.RotateHour = 1
	var err error
	logger, err = rotlog.Rotate()
	if err != nil {
		log.Fatalf("create rotate file error:%s\n", err)
	}

	// defined in test.go for test of use logger
	testRot()
}
