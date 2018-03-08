package daterot

import (
	"time"

	"github.com/sirupsen/logrus"
)

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
)

var (
	// BaseFileName : base file name of log
	BaseFileName = "./log/access.log"
	// BaseLinkName : base link file name of log
	BaseLinkName = ""
	// MaxAgeDays : max days before log file to be purged in file system
	MaxAgeDays time.Duration = 7
	// RotateHour : hour for each rotate
	RotateHour time.Duration = 24
	// LogLevel : log level
	LogLevel = DebugLevel

	// EnableDynamic : enable dynamic modify log level by os Signal
	EnableDynamic = true
	// JSONFormat : use json format by default, false for text
	JSONFormat = true
)
