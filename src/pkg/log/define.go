package log

import (
	"github.com/fatih/color"
	"sync"
)

type Log struct {
	level int
	mu    sync.Mutex
}

const (
	LevelPanic = 0
	LevelError = 1
	LevelWarning = 2
	LevelInfo = 3
	LevelDebug = 4
)

var titles = map[int]string{
	LevelWarning: "Warning",
	LevelPanic:   "Panic",
	LevelError:   "Error",
	LevelInfo:    "Info",
	LevelDebug:   "Debug",
}

// 不同级别前缀与时间的间隔
var spaces = map[int]string{
	LevelWarning: "",
	LevelPanic:   "  ",
	LevelError:   "  ",
	LevelInfo:    "   ",
	LevelDebug:   "  ",
}

var colors = map[int]func(a ...interface{}) string{
	LevelWarning: color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	LevelPanic:   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	LevelError:   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	LevelInfo:    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	LevelDebug:   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

var log *Log
var level = LevelDebug
