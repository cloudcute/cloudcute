package middleware

import (
	"cloudcute/src/pkg/config"
	"cloudcute/src/pkg/log"
	"cloudcute/src/pkg/utils/file_util"
	"cloudcute/src/pkg/utils/path_util"
	"cloudcute/src/routers/api"
	_ "cloudcute/statik" // 嵌入的静态资源, 手动引入后才可直接使用statik包
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"io"
	"net/http"
	"path"
	"strings"
)

// UrlPrefix 访问静态路径前缀
const UrlPrefix = "/"
// 静态中间件跳过处理的路径
var skipPath = []string{ api.UrlPrefix }

type statikFS struct {
	f http.FileSystem
}
func (f *statikFS) Open(name string) (http.File, error) {
	return f.f.Open(name)
}

func (f *statikFS) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if !strings.HasPrefix(p,"/") {
			p = "/" + p;
		}
		if _, err := f.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func initStatic(r *gin.Engine) {
	if !config.SystemConfig.OpenWeb {
		return
	}
	var staticPath = getStaticPath()
	if file_util.Exists(staticPath) {
		// 如果目录存在, 直接使用文件夹中的静态资源
		log.Info("使用本地的静态资源: %s", staticPath)
		r.Use(serve(UrlPrefix, static.LocalFile(staticPath, false)))
		return
	}
	// 使用内嵌的静态资源
	r.Use(serve(UrlPrefix, getStatikFile()))
}

func getStatikFile() *statikFS {
	var f , err = fs.New()
	if err != nil {
		log.Panic("初始化静态资源失败: %s", err)
	}
	return &statikFS{ f : f }
}

func serve(urlPrefix string, fs static.ServeFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		var p = c.Request.URL.Path
		for _ , v := range skipPath {
			if strings.HasPrefix(p, v) {
				c.Next()
				return
			}
		}
		if fs.Exists(urlPrefix, p) || fs.Exists(urlPrefix, path.Join(p, static.INDEX)) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func getStaticPath() string {
	var staticPath string
	if config.IsDev {
		staticPath = path_util.GetAbsPath("./../../public/build")
	}else{
		staticPath = "public"
	}
	return staticPath
}

// ExportStatic 导出静态资源
func ExportStatic() {
	staticFS, err := fs.New()
	if err != nil {
		log.Panic("初始化静态资源失败: %s", err)
	}
	root, err := staticFS.Open("/")
	if err != nil {
		log.Panic("静态资源不存在, %s", err)
	}
	var export func(relPath string, object http.File)
	export = func(relPath string, object http.File) {
		var staticPath = getStaticPath()
		stat, err := object.Stat()
		if err != nil {
			log.Error("读取[%s]的信息失败: %s, 跳过导出...", relPath, err)
			return
		}
		if !stat.IsDir() {
			out, err := file_util.CreatFile(path_util.GetAbsPath(staticPath + relPath))
			defer out.Close()
			if err != nil {
				log.Error("无法创建文件: [%s], %s, 跳过导出...", relPath, err)
				return
			}
			log.Info("导出 [%s]...", relPath)
			if _, err := io.Copy(out, object); err != nil {
				log.Error("无法写入文件[%s], %s, 跳过导出...", relPath, err)
				return
			}
		} else {
			objects, err := object.Readdir(0)
			if err != nil {
				log.Error("无法步入子目录[%s], %s, 跳过导出...", relPath, err)
				return
			}
			for _, newObject := range objects {
				newPath := path.Join(relPath, newObject.Name())
				newRoot, err := staticFS.Open(newPath)
				if err != nil {
					log.Error("无法打开对象[%s], %s, 跳过导出...", newPath, err)
					continue
				}
				export(newPath, newRoot)
			}
		}
	}
	log.Info("开始导出静态资源...")
	export("/", root)
	log.Info("静态资源导出完成")
}
