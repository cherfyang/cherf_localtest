package api

import (
	"awesomeProject1/handler/twilio"
	"github.com/gin-gonic/gin"
)

// 路由绑定 call 函数
func RegisterTwilioRoutes(rt *gin.RouterGroup) {
	rt.POST("/api/voice-call", twilio.Call)
	rt.POST("/process_speech", twilio.ResponseHandle)
}
