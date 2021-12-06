package log

import (
	"fmt"
	"sync"
	"time"
)

type Log struct {
	level int
	mu    sync.Mutex
}

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

var log *Log
var level = LevelDebug

// 不同级别前缀与时间的间隔，保持宽度一致
var spaces = map[string]string{
	"Warning": "",
	"Panic":   "  ",
	"Error":   "  ",
	"Info":    "   ",
	"Debug":   "  ",
}

func GetLog() *Log {
	if log == nil {
		var l = Log{
			level: level,
		}
		log = &l
	}
	return log
}
func SetLevel(level int)  {
	GetLog().level = level
}
// Println 打印行
func (l *Log) Println(prefix string, msg string) {
	// TODO Release时去掉
	// color.NoColor = false
	// c := color.New()
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("%s%s %s %s\n\n","["+prefix+"]", spaces[prefix], time.Now().Format("2006-01-02 15:04:05"), msg)
	//_, _ = c.Printf(
	//	"%s%s %s %s\n",
	//	colors[prefix]("["+prefix+"]"),
	//	spaces[prefix],
	//	time.Now().Format("2006-01-02 15:04:05"),
	//	msg,
	//)
}
// Panic 恐慌
func (l *Log) Panic(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Panic", msg)
	panic(msg)
}
// Error 错误
func (l *Log) Error(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Error", msg)
}
// Warning 警告
func (l *Log) Warning(format string, v ...interface{}) {
	if LevelWarning > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Warning", msg)
}
// Info 信息
func (l *Log) Info(format string, v ...interface{}) {
	if LevelInformational > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Info", msg)
}
// Debug 调试
func (l *Log) Debug(format string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Debug", msg)
}

// Println 打印行
func Println(prefix string, msg string)  {
	GetLog().Println(prefix, msg)
}
// Panic 恐慌
func Panic(format string, v ...interface{})  {
	GetLog().Panic(format, v)
}
// Error 错误
func Error(format string, v ...interface{})  {
	GetLog().Error(format, v)
}
// Warning 警告
func Warning(format string, v ...interface{})  {
	GetLog().Warning(format, v)
}
// Info 信息
func Info(format string, v ...interface{})  {
	GetLog().Info(format, v)
}
// Debug 调试
func Debug(format string, v ...interface{})  {
	GetLog().Debug(format, v)
}
