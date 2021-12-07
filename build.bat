
:: -ldflags "-s   -w"
:: -s去掉符号表（然后panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果）
:: -w去掉DWARF调试信息，得到的程序就不能用gdb调试了

:: -a 强制重新构建

::  -ldflags -X 'main.Version=$VERSION'

@echo off

if "%1"=="-debug" (
    go build -o "build/debug/cloudcute.exe"
) else (
    go build -a -o "build/release/cloudcute.exe" -ldflags "-w -X 'cloudcute/src/module/config.RuntimeMode=release'"
)