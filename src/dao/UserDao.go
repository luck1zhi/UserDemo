package dao

import (
	"UserDemo/src/config"
	"UserDemo/src/model"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

func SelectUserByUserId(userId string) (interface{}, error){
	conn := config.RedisPool.Get()
	defer conn.Close()

	//访问redis
	bytes, err := redis.Bytes(conn.Do("hget","users",userId))
	if err == redis.ErrNil{
		//redis中没有该数据
		return nil, nil
	}
	if err != nil{
		log.Panicf("---> hget(%s)---error:(%v)", userId, err)
	}

	var result model.User
	if err = json.Unmarshal(bytes, &result); err != nil{
		log.Panicf("---> hget(%s) success but Unmarshal failed: %v", userId, err)
	}
	return result,nil
}

func UpdateUser(user model.User) (int,error){

	conn := config.RedisPool.Get()
	defer conn.Close()

	exist, err := redis.Int(conn.Do("hexists", "users", user.UserId))
	if err != nil{
		log.Panicf("---> hexists(%s) failed, error:(%v)", user.UserId, err)
	}
	if exist == 0{
		log.Panicf("---> update(%s) failed because no this user", user.UserId)
	}

	//格式化为JSON
	data, err := json.Marshal(user)
	if err != nil{
		log.Panicf("---> update(%s) failed because Marshal failed: %v", user.UserId, err)
	}

	//更新了返回0
	_, err = conn.Do("hset","users",user.UserId,data)
	if err != nil{
		log.Panicf("---> update(%s) failed, error:(%v)", user.UserId, err)
	}
	return 1, nil
}

func InsertUser(user model.User) (int,error){
	//序列化
	data, err := json.Marshal(user)
	if err != nil{
		log.Panicf("---> insert failed, because Marshal failed: %v", err)
	}

	conn := config.RedisPool.Get()
	defer conn.Close()

	//访问redis
	result, err := redis.Int(conn.Do("hsetnx","users",user.UserId,data))
	if err != nil{
		log.Panicf("---> insert failed, error: %v", err)
	}
	if result == 0{
		log.Panicf("---> insert failed, because user has existed")
	}
	return result, nil
}

func DeleteUser(userId string) (int,error){

	conn := config.RedisPool.Get()
	defer conn.Close()

	//访问redis
	result, err := redis.Int(conn.Do("hdel","users",userId))
	if err != nil{
		log.Panicf("---> delete failed, error: %v", err)
	}
	if result == 0{
		log.Panicf("---> delete failed, because no this user(userId = %s)",userId)
	}
	return result, nil

}

func SelectUsers() (interface{},error){
	conn := config.RedisPool.Get()
	defer conn.Close()

	//获取数据
	data, err := redis.ByteSlices(conn.Do("hvals","users"))
	if err != nil{
		log.Panicf("---> select failed, error: %v", err)
	}

	var results []model.User
	var temp model.User
	for _,val := range data{
		if err = json.Unmarshal(val,&temp); err != nil{
			log.Panicf("---> select success, but Unmarshal failed: %v", err)
		}
		results = append(results, temp)
	}
	return results, nil
}