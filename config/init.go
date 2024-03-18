package config

import "github.com/joho/godotenv"

func InitConfig() error {
	return godotenv.Load()
}
