package config

import (
	"log"
	"os"
)

func GetDataSourceURL() string {
	return GetEnvironmentalValue("DATABASE")

}

func GetPort() string {
	return GetEnvironmentalValue("PORT")

}

func GetEnv() string {
	return GetEnvironmentalValue("ENV")

}

func GetEnvironmentalValue(key string) string {

	if os.Getenv(key) == "" {
		log.Fatalf("%s env var is missing", key)
	}
	return os.Getenv(key)

}
