package router

import (
	"cherf_localtest/router/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		api.RegisterFileRoutes(v1)
		api.RegisterWebRoutes(v1)
	}
}
