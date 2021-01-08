package midware

import (
	"UserDemo/src/common"
	"UserDemo/src/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func MyRecoveryMidware() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err:= recover(); err != nil {
				result := common.ConmmonResult{}
				//链路id
				traceId := c.Writer.Header().Get("X-Request-Trace-Id")
				//堆栈信息
				stackMsg := string(debug.Stack())
				logField := map[string]interface{}{
					"trace_id":    traceId, //  鉴权之后可以得到唯一跟踪ID和用户名
					"user":        c.Writer.Header().Get("X-Request-User"),
					"uri":         c.Request.URL.Path,
					"remote_addr": c.ClientIP(),
					"stack":       stackMsg, // 打印堆栈信息
				}
				c.Abort()
				result.Code, result.Message = 500, fmt.Sprintf("Api内部报错---(id=%s)", traceId)
				//输出到控制台和log文件
				log.Println(logField)
				redisField := make(map[string]interface{})
				for k, v := range logField {
					redisField[k] = v
				}
				redisField["time"] = time.Now().Format("2006-01-02 15:04:05")
				redisField["error"] = err
				dao.AddPanic(redisField) // 上报redis
				c.JSON(http.StatusUnauthorized, result)
				return
			}
		}()
		c.Next()
	}
}