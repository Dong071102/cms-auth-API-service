package main

import (
	"go-cms-backend/config"
	"go-cms-backend/models"
	"go-cms-backend/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

)

func main() {
	godotenv.Load()
	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Lecturer{},
		&models.Admin{},
	)

	e := echo.New()

	// Middleware JWT authentication
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	routes.SetupRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
