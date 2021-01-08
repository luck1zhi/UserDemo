package service

import (
	"UserDemo/src/dao"
	"UserDemo/src/model"
)

func AddUser(user model.User) (int,error){
	return dao.InsertUser(user)
}

func DelUser(userId string) (int,error){
	return dao.DeleteUser(userId)
}

func GetUser(userId string) (interface{},error){
	return dao.SelectUserByUserId(userId)
}

func ModifyUser(user model.User) (int,error){
	return dao.UpdateUser(user)
}

func GetUserList() (interface{},error){
	return dao.SelectUsers()
}