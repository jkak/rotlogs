package sizerot

import "github.com/sirupsen/logrus"

// log level
const (
	// PanicLevel panic
	PanicLevel = logrus.PanicLevel
	// FatalLevel fatal
	FatalLevel = logrus.FatalLevel
	// ErrorLevel error
	ErrorLevel = logrus.ErrorLevel
	// WarnLevel warning
	WarnLevel = logrus.WarnLevel
	// InfoLevel info
	InfoLevel = logrus.InfoLevel
	// DebugLevel debug
	DebugLevel = logrus.DebugLevel

	fmtStr = "2006-01-02 15:04:05"
)

var (
	// BaseFileName : base file name of log
	BaseFileName = "./log/access.log"

	// MaxMegaSize : the maximum size in megabytes of the log file before it gets rotated
	MaxMegaSize = 100
	// MaxBackups : the maximum number of old log files to retain
	MaxBackups = 7
	// Compress : compress by default.
	Compress = true
	// MaxAgeDays : max days before log file to be purged in file system
	MaxAgeDays = 7
	// LogLevel : log level
	LogLevel = DebugLevel

	// EnableDynamic : enable dynamic modify log level by os Signal
	EnableDynamic = false
	// JSONFormat : use json format by default, false for text
	JSONFormat = true
)
