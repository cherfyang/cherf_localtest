package file_handle

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

func DeleteFile(c *gin.Context) {
	path := c.Param("path")
	token := c.Param("token")

}
