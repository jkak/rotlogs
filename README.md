## rotlogs is a golang log package

rotlogs是基于其他开源日志包封装而成的，可进行日志轮转的golang日志处理包。按场景，实现了两种日志轮转方式：

* 按日期轮转。
  * 基于logrus和file-rotatelogs
* 按文件大小进行轮转。
  * 基于logrus和lumberjack



### 一、日期轮转

#### 常量

```go
// 6个日志级别：
PanicLevel
FatalLevel
ErrorLevel
WarnLevel
InfoLevel
DebugLevel
```

#### 变量

```go
// 相关变量主要用于变量日志的默认行为，所有变量都有默认值 
BaseFileName	// 日志文件名，默认./log/access.log
BaseLinkName	// 日志软链名，默认为空，为空时则没有软链
MaxAgeDays		// 日志保留天数，默认为7天
RotateHour		// 日志切割间隔，默认24小时
LogLevel		// 日志级别，默认为DebugLevel
JSONFormat		// 日志格式，json或text。默认为true，使用json格式。

EnableDynamic	// 允许动态调整日志级别。默认true
```

#### 使用示例

```go
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
		time.Sleep(500 * time.Millisecond)
	}
}
```

#### 动态调整日志级别

```shell
# 开启动态调整时，使用信号来控制：
# 	USR1 to turn down log level from Panic to Debug
# 	USR2 to turn - up log level from Debug to Panic
# usage: 

kill -s USR1|USR2  XX_PID
```



### 二、大小轮转





