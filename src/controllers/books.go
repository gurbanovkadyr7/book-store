package controllers

import (
	ser "admin/src/services"

	"github.com/gin-gonic/gin"
)

func CreateBooks(router *gin.Engine) {
	group := router.Group("books")

	group.GET("/all", ser.Get_books)
	group.POST("/create", ser.Create_books)
	group.POST("/by_book_name", ser.Get_by_bookName)
	group.POST("/by_author", ser.Get_by_author)
	group.POST("/by_category", ser.Get_by_category)
}
