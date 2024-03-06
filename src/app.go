package src

import (
	"admin/config"
	"admin/src/controllers"
	"admin/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Init_app() *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies(nil)

	router.Use(gin.Logger())
	router.Use(middlewares.Cors)
	router.Static("/public", config.ENV.DIR)

	controllers.CreateUser(router)
	controllers.CreateBooks(router)

	return router
}
