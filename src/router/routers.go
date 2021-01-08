package router

import (
	"UserDemo/src/config"
	"UserDemo/src/controller"
	"UserDemo/src/midware"
	"UserDemo/src/model"
	"UserDemo/src/validate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegistryRoute(r *gin.Engine){

	//中间件的导入需要在注册路由之前
	r.Use(gin.LoggerWithFormatter(config.MyLogFormatter))
	r.Use(midware.MyRecoveryMidware())

	//一定要在添加中间件之后！否则不会生效
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("checkUserId", validate.CheckUserId)
		v.RegisterValidation("checkUserName", validate.CheckUserName)
	}

	r.GET("/ping",controller.Ping)
	r.GET("/jsondemo", func(context *gin.Context) {
		context.JSON(http.StatusOK,model.User{UserId: "1111",NickName: "zcz",Role: 1,Age: 16})
	})

	userG := r.Group("/user")

	userG.POST("/add_user",controller.AddUser)
	userG.DELETE("/del_user/:id", controller.DelUser)
	userG.GET("/get_user/:id", controller.GetUser)
	userG.PUT("/modify_user", controller.ModifyUser)
	userG.GET("/get_user_list", controller.GetUserList)
}