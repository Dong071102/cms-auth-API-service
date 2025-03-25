package main

import (
	"fmt"
	"auth-cms-backend/config"
	"auth-cms-backend/models"
	"auth-cms-backend/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	fmt.Println("Loaded JWT_SECRET:", os.Getenv("JWT_SECRET"))
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
