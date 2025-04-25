package log

import (
	_const "cherf_localtest/const"
	"cherf_localtest/sdk"
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
	fd, err := os.OpenFile(todir+"upload.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer fd.Close()
	defer f.Close()
	speed := float64(fileSize) / usetime // bytes per second
	speedStr := fmt.Sprintf("%.2f MB/s", speed/1024/1024)

	logEntry := fmt.Sprintf(
		"[%s]\n"+
			"\t[文件名]      : %s\n"+
			"\t[文件大小]    : %s\n"+
			"\t[上传来源]    : %s\n"+
			"\t[上传目标]    : %s\n"+
			"\t[消耗时间]        : %.3f\t\t\t仅作参考\n"+
			"\t[传输速度]    : %s\n"+
			"\t来自adress :%s\n"+
			"\tUser-Agent : %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		fileName,
		util.FileSizeConvert(fileSize),
		from,
		todir,
		usetime,
		speedStr,
		sdk.IpToAddress(clientIP),
		Agent,
	)
	fd.WriteString(logEntry)
	f.WriteString(logEntry)
}

func LogDownload(c *gin.Context, start time.Time, usetime float64) {
	f, err := os.OpenFile(_const.ApiLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	util.DebugRequest(c)
	logEntry := ""

	logEntry += fmt.Sprintf("ip", c.ClientIP())

	f.WriteString(logEntry)
}

func Info(info string) {
	f, err := os.OpenFile(_const.InfoLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	info = "------------------------日志级别:infomation------------------------\n" + info + "------------------------日志级别:infomation------------------------"
	f.WriteString(info)
}
func Err(info string) {
	f, err := os.OpenFile(_const.ErrLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	info = "------------------------日志级别:error------------------------\n" + info + "------------------------日志级别:error------------------------\n"
	f.WriteString(info)
}
func RequestInfo(c *gin.Context) {
	f, err := os.OpenFile(_const.RequestLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	req := util.RequestString(c)
	info := "------------------------请求信息------------------------\n" + req + "------------------------请求信息------------------------\n"
	f.WriteString(info)

}
