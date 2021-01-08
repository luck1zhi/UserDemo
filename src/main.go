package main

import (
	"UserDemo/src/config"
	"UserDemo/src/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init(){
	//配置redis连接池
	config.RedisPool = config.NewRedisPool()
}

func fileExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	//r := gin.Default()
	r := gin.New()

	//配置日志格式、输出
	var logFile *os.File
	if fileExist(config.LogFileLocation){
		logFile,_ = os.OpenFile(config.LogFileLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	}else{
		logFile,_ = os.Create(config.LogFileLocation)
	}
	log.SetOutput(io.MultiWriter(logFile,os.Stdout))
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	defer logFile.Close()


	router.RegistryRoute(r)

	r.Run(":9999")
	//srv := &http.Server{
	//	Addr:    ":9999",
	//	Handler: r,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//srv.ListenAndServe()


}
