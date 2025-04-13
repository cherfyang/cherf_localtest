package log

import (
	_const "cherf_localtest/const"
	"cherf_localtest/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func LogUpload(clientIP, fileName string, fileSize int64, todir string, Agent string, from string, usetime float64) {

	f, err := os.OpenFile(_const.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	speed := float64(fileSize) / usetime // bytes per second
	speedStr := fmt.Sprintf("%.2f MB/s", speed/1024/1024)

	logEntry := fmt.Sprintf(
		"[%s]\n"+
			"\t文件名      : %s\n"+
			"\t文件大小    : %s\n"+
			"\t上传来源    : %s\n"+
			"\t上传目标    : %s\n"+
			"\t耗时        : %.3f\n"+
			"\t传输速度    : %s\n"+
			"\t来自 IP     : %s\n"+
			"\tUser-Agent : %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		fileName,
		util.FileSizeConvert(fileSize),
		from,
		todir,
		usetime,
		speedStr,
		clientIP,
		Agent,
	)

	f.WriteString(logEntry)
}

func LogDownload(c *gin.Context) {
	f, err := os.OpenFile(_const.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	util.DebugRequest(c)

	logEntry := fmt.Sprintf("" +
		"")

	f.WriteString(logEntry)
}
