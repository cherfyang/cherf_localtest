package api

import (
	"cherf_localtest/handler/file_handle"
	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(rf *gin.RouterGroup) {
	//param 附带要分类的文件路径
	rf.POST("/file/category", file_handle.FileCategory)
	//body 里写要发送给哪个 url,以及发送的文件路径
	rf.POST("/file/sendfile", file_handle.SendFileHandle)

	rf.POST("/file/categorybyname", file_handle.FileCategoryByFileNme)
	rf.POST("/file/upload", file_handle.UploadHandler)
	rf.GET("/file/list", file_handle.ListHandler)
	rf.GET("/file/download", file_handle.DownloadFile)
	rf.GET("/file/content", file_handle.LoadFile)
	rf.GET("/video/play", file_handle.StreamVideo)
	rf.POST("/file/save", file_handle.SaveFileContent)
}
