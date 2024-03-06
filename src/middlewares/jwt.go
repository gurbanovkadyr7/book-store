package middlewares

import (
	"admin/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	First_name string
	Hash_pass  string
}

func Guard(ctx *gin.Context) {
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
			return []byte(config.ENV.ACCESS_KEY), nil
		},
	)

	if err != nil {
		ctx.AbortWithStatus(403)
		return
	}

	user := User{
		First_name: string(claims["first_name"].(string)),
		Hash_pass:  string(claims["hash_pass"].(string)),
	}

	ctx.Set("users", user)
	ctx.Next()
}
