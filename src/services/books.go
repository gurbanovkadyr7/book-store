package services

import (
	"admin/src/repositories"
	rep "admin/src/repositories"
	schemas "admin/src/schemas/request"
	response "admin/src/schemas/response"

	"github.com/gin-gonic/gin"
)

func Get_books(ctx *gin.Context) {
	books, err := repositories.Get_book()

	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.IndentedJSON(200, books)
}

func Create_books(ctx *gin.Context) {
	var book schemas.Create_books

	validate_err := ctx.BindJSON(&book)

	if validate_err != nil {
		ctx.JSON(400, validate_err.Error())
	}

	err := rep.Create_book(book)

	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.JSON(200, gin.H{"success": "created"})

}

func Get_by_bookName(ctx *gin.Context) {
	var bookName response.Book_ByName

	validate_err := ctx.BindJSON(&bookName)

	if validate_err != nil {
		ctx.JSON(400, validate_err)
	}

	book, err := rep.BookByName(bookName)

	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.IndentedJSON(200, book)
}

func Get_by_author(ctx *gin.Context) {
	var bookAuthor response.Book_ByAuthor

	validate_err := ctx.BindJSON(&bookAuthor)

	if validate_err != nil {
		ctx.JSON(400, validate_err)
	}

	book, err := rep.BookByAuthor(bookAuthor)

	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.IndentedJSON(200, book)
}

func Get_by_category(ctx *gin.Context) {
	var category response.Book_ByCategory

	validate_err := ctx.BindJSON(&category)

	if validate_err != nil {
		ctx.JSON(400, validate_err.Error())
	}

	books, err := rep.BookByCategory(category)

	if err != nil {
		ctx.JSON(400, err.Error())
	}

	ctx.IndentedJSON(200, books)
}
