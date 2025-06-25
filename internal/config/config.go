package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBUrl      string
	Port       string
	SecretKey  string
	ApiKey     string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("user")
	password := url.QueryEscape(os.Getenv("password"))
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	return Config{
		DBUser:     user,
		DBPassword: password,
		DBHost:     host,
		DBPort:     port,
		DBName:     dbname,
		DBUrl:      dbUrl,
		Port:       os.Getenv("port"),
		SecretKey:  os.Getenv("secret_key"),
		ApiKey:     os.Getenv("API_KEY"),
	}
}
