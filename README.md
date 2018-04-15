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

LoggerPtr // 提供一个日志句柄变量，可以用户的main之外定义，在其他文件中统一使用
```

#### 使用示例

参考：

* example/date/date.go

#### 动态调整日志级别

```shell
# 开启动态调整时，使用信号来控制：
# 	USR1 to turn down log level from Panic to Debug
# 	USR2 to turn - up log level from Debug to Panic
# usage: 

kill -s USR1|USR2  XX_PID
```



### 二、大小轮转

#### 常量

常量同日期轮转。

#### 变量

```go
// 相关变量主要用于变量日志的默认行为，所有变量都有默认值
BaseFileName	// 日志文件名，默认./log/access.log
MaxMegaSize		// 日志文件大小，超过后将轮转，默认100MB
MaxBackups		// 日志保留份数，默认7个
MaxAgeDays		// 日志保留天数，默认7天
Compress		// 日志压缩，默认为true
LogLevel		// 日志级别，默认为DebugLevel
JSONFormat		// 日志格式，json或text。默认为true，使用json格式。

EnableDynamic	// 允许动态调整日志级别。默认true
```

#### 使用示例

参考：

- example/size/size.go


### 参考

本项目主要参考[bigwhite](https://github.com/bigwhite)的[博客](https://tonybai.com/2018/01/13/the-problems-i-encountered-when-writing-go-code-issue-1st/)。




