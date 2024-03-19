package main

import (
	"fmt"
	"fremont/config"
	postgresqlpkg "fremont/pkg/database/postgresql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := config.InitConfig(); err != nil {
		panic("error init config")
	}
	router := SetupRouter()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	log.Fatal(router.Start(port))

}

func SetupRouter() *echo.Echo {
	// setup database connection
	log.Println("Database Connection")
	_ = postgresqlpkg.InitPostgreSQL()

	// setup echo
	e := echo.New()

	// setup CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))
	//init router

	fmt.Println(e)
	return e
}
