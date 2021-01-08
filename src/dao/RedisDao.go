package dao

import (
	"UserDemo/src/config"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

/**
上报异常
 */
func AddPanic(logInfo map[string]interface{}) error{

	defer func() {
		//...
		recover()
	}()

	//获取Redis连接
	conn := config.RedisPool.Get()
	defer conn.Close()

	//map-log序列化
	data,err := json.Marshal(logInfo)
	if err != nil {
		return err
	}

	//访问redis
	//trace_id为key
	_, err = redis.Int(conn.Do("hset","panic", logInfo["trace_id"], data))
	if err != nil {
		return err
	}
	return nil
}