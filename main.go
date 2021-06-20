package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func toLinux(basePath string) string {
	return strings.ReplaceAll(basePath, "\\", "/")
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
func createMyRender(dirname string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
		filename := toLinux(path[len(dirname)-1:])
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			filename = strings.Trim(filename,"/")
			filename = strings.Trim(filename,"\\")
			fmt.Println(filename,path)
			r.AddFromFiles(filename, path)
		}
		return nil
	})

	return r
}

func main() {
	var publicPath string
	var staticPath string
	var outport int

	flag.StringVar(&publicPath, "w", "./website", "静态文件地址")
	flag.StringVar(&staticPath, "s", "./website/static", "静态资源地址")
	flag.IntVar(&outport, "o", 8044, "输出端口")
	if !flag.Parsed() {
		flag.Parse()
	}

	engine := gin.Default()
	engine.Static("/static", staticPath)
	iconPath := path.Join(publicPath, "favicon.ico")
	if IsExist(iconPath) {
		engine.Use(favicon.New(iconPath))
	} else {
		iconPath = path.Join(".", "favicon.icon")
		if IsExist(iconPath) {
			engine.Use(favicon.New(iconPath))
		}
	}

	engine.HTMLRender = createMyRender(publicPath)
	engine.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{})
	})
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{})
	})

	filepath.Walk(publicPath, func(path string, info fs.FileInfo, err error) error {
		filename := toLinux(path[len(publicPath)-1:])
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			engine.GET(filename, func(c *gin.Context) {
				c.HTML(200, filename, gin.H{})
			})
		}

		return nil
	})

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{})
	})

	outputStr := strconv.Itoa(outport)
	fmt.Println("打开浏览器：http://localhost:" + outputStr)
	engine.Run(":" + outputStr)
}
