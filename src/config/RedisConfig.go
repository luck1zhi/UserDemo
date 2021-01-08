package config

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var(
	RedisPool *redis.Pool
	redisHost="localhost:6379"
	redisDB = redis.DialDatabase(0)
	redisPass = redis.DialPassword("root")
)

func NewRedisPool()*redis.Pool{
	return &redis.Pool{
		MaxIdle:50,         //最大空闲连接数
		MaxActive:30,        //允许分配最大连接数
		IdleTimeout:300*time.Second,    //连接时间限制
		Dial: func() (redis.Conn,  error) {    //创建连接的函数
			redisConn,err := redis.Dial("tcp",redisHost,redisDB,redisPass)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			return redisConn,nil
		},
	}
}
