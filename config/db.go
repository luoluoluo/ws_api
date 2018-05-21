package config

import "os"

// Db config db
var DB = map[string]string{
	"user":     os.Getenv("DB_USER"),
	"password": os.Getenv("DB_PASSWORD"),
	"dbname":   os.Getenv("DB_NAME"),
	"host":     os.Getenv("DB_HOST"),
	"port":     os.Getenv("DB_PORT"),
}
