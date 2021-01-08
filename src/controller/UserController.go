package controller

import (
	"UserDemo/src/common"
	"UserDemo/src/model"
	"UserDemo/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK,"tong")
}

func AddUser(c *gin.Context) {
	var user model.User
	if err:=c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 777, Message: err.Error()})
	}else{
		data, err := service.AddUser(user)
		if err != nil{
			c.JSON(http.StatusOK,common.ConmmonResult{Code: 416, Message: err.Error()})
		}else{
			c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "添加成功...", Data: data})
		}
	}
}

func DelUser(c *gin.Context){
	userId := c.Param("id")
	c.ShouldBindJSON(userId)

	data, err := service.DelUser(userId)
	if err != nil{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 416, Message: err.Error()})
	}else{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "删除成功...", Data: data})
	}
}

func GetUser(c *gin.Context){
	userId := c.Param("id")
	c.ShouldBindJSON(userId)
	data, err := service.GetUser(userId)

	if err != nil{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "查询成功...", Data: data})
	}else{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "没有该用户信息...", Data: data})
	}
}

func ModifyUser(c *gin.Context){
	var user model.User
	c.ShouldBindJSON(&user)

	_, err := service.ModifyUser(user)
	if err != nil{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 416, Message: err.Error()})
	}else{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "更新成功...", Data: user})
	}

}

func GetUserList(c *gin.Context){
	data, err := service.GetUserList()
	if err != nil{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 416, Message: err.Error()})
	}else{
		c.JSON(http.StatusOK,common.ConmmonResult{Code: 200, Message: "查询成功...", Data: data})
	}
}