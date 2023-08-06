package main

import (
	"mini-tiktok/controller"
	"mini-tiktok/middleware"
	"mini-tiktok/model"

	"github.com/gin-gonic/gin"
)

func main() {
	err := model.InitMysql()
	if err != nil {
		panic("DB wrong")
	}
	r := gin.Default()
	r.POST("/user/login", middleware.MD5Middleware(), controller.UserLoginHandler)
	r.POST("/user/register", middleware.MD5Middleware(), controller.UserRegisterHandler)
	r.Run(":8080")
}