package main

import (
	"awesomeProject1/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.RegisterRoutes(r)

	r.Run(":8080")
}
