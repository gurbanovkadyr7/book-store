package services

import (
	"admin/config"
	rep "admin/src/repositories"
	schemas "admin/src/schemas/request"
	response "admin/src/schemas/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Get_users(ctx *gin.Context) {
	users, err := rep.Get()

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.IndentedJSON(200, users)
}

func Create_users(ctx *gin.Context) {
	var user schemas.Create_users

	validate_err := ctx.BindJSON(&user)

	if validate_err != nil {
		ctx.JSON(400, validate_err.Error())
		return
	}

	err := rep.Create_user(user)

	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(201, gin.H{"message": "Successfully created"})
}

func Login_users(ctx *gin.Context) {
	var user response.User

	validate_err := ctx.BindJSON(&user)

	if validate_err != nil {
		ctx.JSON(500, validate_err.Error())
	}

	users, err := rep.Check_User(user)

	if err != nil {
		ctx.JSON(500, err.Error())
	}

	token := rep.Create_token(users)

	ctx.IndentedJSON(200, token)

}

func Refresh(ctx *gin.Context) {
	authorization := ctx.Request.Header["Authorization"]

	if len(authorization) == 0 {
		ctx.AbortWithStatus(401)
		return
	}

	split_bearer := strings.Split(authorization[0], "Bearer ")

	if len(split_bearer) == 0 {
		ctx.AbortWithStatus(401)
		return
	}

	token := split_bearer[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.ENV.REFRESH_KEY), nil
		},
	)

	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}

	new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"first_name": string(claims["first_name"].(string)),
		"hash_pass":  string(claims["hash_pass"].(string)),
		"exp":        time.Now().Add(config.ENV.ACCESS_TIME).Unix(),
	})

	token_string, _ := new_token.SignedString([]byte(config.ENV.ACCESS_KEY))

	ctx.JSON(200, token_string)
}
