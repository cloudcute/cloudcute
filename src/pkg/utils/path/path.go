package path

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// RemoveSlash 移除路径最后的`/`
func RemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// Split 分割路径为列表
func Split(path string) []string {
	if len(path) == 0 || path[0] != '/' {
		return []string{}
	}
	if path == "/" {
		return []string{"/"}
	}
	pathSplit := strings.Split(path, "/")
	pathSplit[0] = "/"
	return pathSplit
}

func Join(elem ...string) string {
	return path.Join(elem...)
}

// GetAbsPath 从相对可执行文件的路径转绝对路径
func GetAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), path)
}
