package sizerot

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// Rotate for rotate log file
func Rotate() *logrus.Logger {
	// rotate and compress writter
	sizeLog := &lumberjack.Logger{
		Filename:   BaseFileName,
		MaxSize:    MaxMegaSize, // megabytes
		MaxBackups: MaxBackups,  // file numbers
		MaxAge:     MaxAgeDays,  // days
		Compress:   Compress,    // disabled by default
		LocalTime:  true,
	}

	// logger
	logger := logrus.New()
	logger.SetLevel(LogLevel)
	if JSONFormat {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{}
	}
	// dynamic modify log level by signal
	if EnableDynamic {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2)
		go sigModLoglevel(sig, logger)
	}

	// logger interface wrapper
	logger.Out = sizeLog
	return logger
}

// accept signal to modify log level dynamicall
// USR1 to turn down log level from Panic to Debug
// USR2 to turn - up log level from Debug to Panic
func sigModLoglevel(sigCh chan os.Signal, logger *logrus.Logger) {
	for {
		select {
		case sig := <-sigCh:
			modBySig(sig, logger)
		}
	}
}
func modBySig(sig os.Signal, logger *logrus.Logger) {
	if sig == syscall.SIGUSR1 {
		level := logger.Level
		if level != PanicLevel {
			logger.SetLevel(level - 1)
		}
		logrus.Println(time.Now().Format(fmtStr), "Raise log level to:", logger.Level)

	} else if sig == syscall.SIGUSR2 {
		level := logger.Level
		if level != DebugLevel {
			logger.SetLevel(level + 1)
		}
		logrus.Println(time.Now().Format(fmtStr), "Reduce log level to:", logger.Level)
	} else {
		logrus.Println(time.Now().Format(fmtStr), "receive unknown signal:", sig)
	}
}
