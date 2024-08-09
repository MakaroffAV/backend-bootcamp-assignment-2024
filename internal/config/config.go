package config

import "fmt"

var DbTplString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow"

var DbParams = map[string]string{
	"host": "localhost",
	"port": "20500",
	"user": "makarov",
	"pass": "simplePassword",
	"name": "bootcamp",
}

var DbConnString = fmt.Sprintf(
	DbTplString,
	DbParams["host"],
	DbParams["port"],
	DbParams["user"],
	DbParams["pass"],
	DbParams["name"],
)
