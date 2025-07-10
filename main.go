package main

import (
	"cherf_localtest/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("html/*")

	r.StaticFile("/favicon.ico", "D:\\code\\cherf_localtest\\static\\favicon.ico") // 指定文件路径

	//r.Static("/static", "./static")

	// 页面路由
	r.GET("/file", func(c *gin.Context) {
		c.HTML(200, "文件管理页面.html", nil)
	})
	r.GET("/file2", func(c *gin.Context) {
		c.HTML(200, "文件目录.html", nil)
	})
	r.GET("/filelist/:namepath", func(c *gin.Context) {
		c.HTML(200, "文件列表.html", nil)
	})
	r.GET("/editor.html", func(c *gin.Context) {
		c.HTML(200, "editor.html", nil)
	})
	r.GET("/website", func(c *gin.Context) {
		c.HTML(200, "web_navigation.html", nil)
	})
	// 注册 API 路由
	router.RegisterRoutes(r)

	r.Run("0.0.0.0:8080")
}
