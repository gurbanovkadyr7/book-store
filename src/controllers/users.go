package controllers

import (
	"admin/src/middlewares"
	ser "admin/src/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(router *gin.Engine) {
	group := router.Group("users")

	group.GET("/all", ser.Get_users)
	group.POST("/create", middlewares.Guard, ser.Create_users)
	group.POST("/login", ser.Login_users)
	group.GET("/refresh", ser.Refresh)
}


