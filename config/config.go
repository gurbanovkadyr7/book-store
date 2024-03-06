package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER string
	DB_PORT string
	DB_HOST string
	DB_NAME string
	DB_PASS string

	API_HOST     string
	API_PORT     string
	ACCESS_TIME  time.Duration
	REFRESH_TIME time.Duration
	ACCESS_KEY   string
	REFRESH_KEY  string
	DIR          string
}

var ENV Config

func Init_Config() {
	godotenv.Load()

	ENV.DB_USER = os.Getenv("DB_USER")
	ENV.DB_PORT = os.Getenv("DB_PORT")
	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.DB_PASS = os.Getenv("DB_PASS")

	ENV.API_HOST = os.Getenv("API_HOST")
	ENV.API_PORT = os.Getenv("API_PORT")
	AT, _ := time.ParseDuration(os.Getenv(("ACCESS_TIME")))
	ENV.ACCESS_TIME = AT
	RT, _ := time.ParseDuration(os.Getenv("REFRESH_TIME"))
	ENV.REFRESH_TIME = RT
	ENV.ACCESS_KEY = os.Getenv("ACCESS_KEY")
	ENV.REFRESH_KEY = os.Getenv("REFRESH_KEY")
	ENV.DIR = os.Getenv("DIR")
}
