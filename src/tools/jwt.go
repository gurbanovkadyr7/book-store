package tools

import "github.com/golang-jwt/jwt/v5"

func Create_token(first_name string, hash_pass string, time int64, secret_key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"first_name": first_name,
		"hash_pass":  hash_pass,
		"exp":        time,
	})

	token_string, _ := token.SignedString([]byte(secret_key))

	return token_string
}
