package log

import (
	"cloudcute/src/pkg/utils/utils"
	"fmt"
	"github.com/fatih/color"
	"time"
)

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
func (l *Log) printColor(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	var colorLog = color.New()
	_, _ = colorLog.Printf(msg)
}
func formatStr(format string, v ...interface{}) string {
	var hasValue = v != nil && len(v) > 0
	if hasValue {
		return fmt.Sprintf(format, v...)
	}else{
		return format
	}
}
// Println 打印行
func (l *Log) Println(title string, msg string) {
	msg = formatStr("%s %s %s\n", title, time.Now().Format("2006-01-02 15:04:05"), msg)
	l.printColor(msg)
}
func (l *Log) PrintLevel(level int, format string, v ...interface{}) {
	if level > l.level {
		return
	}
	var msg = formatStr(format, v...)
	var titleStr = colors[level](formatStr("[%s]", titles[level]))
	var title = formatStr("%s%s", titleStr, spaces[level])
	l.Println(title, msg)
	if level == LevelPanic {
		utils.WaitExit()
	}
}

// Panic 恐慌
func (l *Log) Panic(format string, v ...interface{}) {
	l.PrintLevel(LevelPanic, format, v...)
}
// Error 错误
func (l *Log) Error(format string, v ...interface{}) {
	l.PrintLevel(LevelError, format, v...)
}
// Warning 警告
func (l *Log) Warning(format string, v ...interface{}) {
	l.PrintLevel(LevelWarning, format, v...)
}
// Info 信息
func (l *Log) Info(format string, v ...interface{}) {
	l.PrintLevel(LevelInfo, format, v...)
}
// Debug 调试
func (l *Log) Debug(format string, v ...interface{}) {
	l.PrintLevel(LevelDebug, format, v...)
}

// Println 打印行
func Println(title string, msg string)  {
	GetLog().Println(title, msg)
}
// PrintLevel 打印等级
func PrintLevel(level int, msg string)  {
	GetLog().PrintLevel(level, msg)
}
// Panic 恐慌
func Panic(format string, v ...interface{})  {
	GetLog().Panic(format, v...)
}
// Error 错误
func Error(format string, v ...interface{})  {
	GetLog().Error(format, v...)
}
// Warning 警告
func Warning(format string, v ...interface{})  {
	GetLog().Warning(format, v...)
}
// Info 信息
func Info(format string, v ...interface{})  {
	GetLog().Info(format, v...)
}
// Debug 调试
func Debug(format string, v ...interface{})  {
	GetLog().Debug(format, v...)
}
