package config

import "os"

var Db = map[string]string{
	"user":     os.Getenv("DB_USER"),
	"password": os.Getenv("DB_PASSWORD"),
	"dbname":   os.Getenv("DB_NAME"),
	"host":     os.Getenv("DB_HOST"),
	"port":     os.Getenv("DB_PORT"),
}
