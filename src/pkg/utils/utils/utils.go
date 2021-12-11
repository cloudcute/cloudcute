package utils

import (
	"fmt"
	"os"
)

// Pause 程序暂停 按键继续
func Pause() {
	fmt.Println("请按Enter键继续...")
	fmt.Scanln()
//	fmt.Println("请按任意键继续...")
//Loop:
//	for {
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			break Loop
//		}
//	}
}
// WaitExit 等待按键并退出程序
func WaitExit() {
	Pause()
	Exit()
}
// Exit 退出程序
func Exit() {
	ExitCode(0)
}
// ExitError 错误退出
func ExitError() {
	ExitCode(1)
}
// ExitCode 指定Code退出程序
func ExitCode(code int) {
	os.Exit(code)
}
