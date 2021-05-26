package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)


func toLinux(basePath string) string {
	return strings.ReplaceAll(basePath, "\\", "/")
}

func createMyRender(dirname string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
		filename := toLinux(path[len(dirname)-1:])
		if !info.IsDir() && strings.HasSuffix(path,".html") {
			r.AddFromFiles(filename,path)
		}
		return nil
	})


	return r
}

func main() {
	var websitePath string
	var staticPath string
	var outport int

	flag.StringVar(&websitePath,"w","./website","静态文件地址")
	flag.StringVar(&staticPath,"s","./public","静态资源地址")
	flag.IntVar(&outport,"o",8044,"输出端口")
	if !flag.Parsed(){
		flag.Parse()
	}

	fmt.Println("ok" )
	engine := gin.Default()
	engine.Static("/static", staticPath)
	fmt.Println(engine)
	engine.HTMLRender = createMyRender(websitePath)
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index.html")
	})
	filepath.Walk(websitePath, func(path string, info fs.FileInfo, err error) error {
		filename := toLinux(path[len(websitePath)-1:])
		if !info.IsDir() && strings.HasSuffix(path,".html") {
			engine.GET(filename, func(c *gin.Context) {
				c.HTML(200, filename, gin.H{})
			})
		}

		return nil
	})

	outputStr := strconv.Itoa(outport)
	engine.Run(":"+ outputStr)
}
