package main

import (
	"admin/config"
	"admin/database"
	app "admin/src"
	"fmt"
)

func main() {
	config.Init_Config()
	database.Init_db()
	server := app.Init_app()
	address := fmt.Sprintf("%v:%v", config.ENV.API_HOST, config.ENV.API_PORT)
	server.Run(address)
}
