package routes

import (
	"go-cms-backend/controllers"
	custommiddleware "go-cms-backend/middleware"

	"github.com/labstack/echo/v4"
)
func SetupRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "API is running ✅"})
	})
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)

	// Protected route example:
	e.GET("/me", func(c echo.Context) error {
		claims := c.Get("user")
		return c.JSON(200, map[string]any{"message": "You are authenticated", "claims": claims})
	}, custommiddleware.JWTAuthMiddleware)

	// Admin-only route example:
	e.GET("/admin-only", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome, admin!"})
	}, custommiddleware.JWTAuthMiddleware, custommiddleware.RoleMiddleware("admin"))

	// Student-only route example:
	e.GET("/student-only", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome, student!"})
	}, custommiddleware.JWTAuthMiddleware, custommiddleware.RoleMiddleware("student"))

	// Lecturer-only route example:
	e.GET("/lecturer-only", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome, lecturer!"})
	}, custommiddleware.JWTAuthMiddleware, custommiddleware.RoleMiddleware("lecturer"))
}
