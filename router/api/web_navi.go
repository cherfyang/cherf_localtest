package api

import (
	handlers "cherf_localtest/handler/webnavigation"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(rw *gin.RouterGroup) {
	rw.GET("/webs/list", handlers.GetWebs)
	rw.DELETE("/webs/delete", handlers.DeleteWeb)
	rw.POST("/webs/add", handlers.AddWeb)
	rw.POST("/webs/undo-delete", handlers.UndoDeleteWeb)
}
