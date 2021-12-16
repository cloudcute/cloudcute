
# -ldflags "-s   -w"
# -s去掉符号表（然后panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果）
# -w去掉DWARF调试信息，得到的程序就不能用gdb调试了
# -a 强制重新构建
#  -ldflags -X 'main.Version=$VERSION'

if [ "$1"x = "-debug"x ]; then
    go build -o "build/debug/cloudcute"
else
    statik -f -src=public/build/ -include=*.html,*.js,*.json,*.css,*.png,*.jpg,*.svg,*.ico,*.ttf,*.map,*.txt
    go build -a -o "build/release/cloudcute" -ldflags "-w -X 'cloudcute/src/pkg/config.RuntimeMode=release'"
fi