package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	LogFileLocation string = "userdemo.log"
)

func MyLogFormatter(param gin.LogFormatterParams) string {
	// 你的自定义格式
	return fmt.Sprintf("[%s] : [%d] -- %s -- (%s) -- |%s| -- %s -- Error=\"%s\"\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.StatusCode,
		param.ClientIP,
		param.Method,
		param.Path,
		param.Latency,
		param.ErrorMessage,
	)
}