package config

import "os"

func ServerPort() string {
	if os.Getenv("APP_PORT") == "" {
		return ":8080"
	}
	return ":" + os.Getenv("APP_PORT")
}

func ServerEnvirontment() string {
	if os.Getenv("APP_ENV") == "" {
		return "local"
	}
	return os.Getenv("APP_ENV")
}
