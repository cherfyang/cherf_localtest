package router

import (
	"awesomeProject1/router/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		api.RegisterTwilioRoutes(v1)
		api.RegisterFileRoutes(v1)
	}
}
