package main

import (
	"cherf_localtest/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.RegisterRoutes(r)

	r.Run("0.0.0.0:8080")
}
