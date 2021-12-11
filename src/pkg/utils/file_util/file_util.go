package file_util

import (
	"cloudcute/src/pkg/log"
	"io"
	"os"
	"path/filepath"
)

// Exists 目录或文件是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreatFile 给定path创建文件，如果目录不存在就递归创建
func CreatFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			log.Warning("无法创建目录，%s", err)
			return nil, err
		}
	}
	return os.Create(path)
}

// IsEmpty 返回给定目录是否为空目录
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
