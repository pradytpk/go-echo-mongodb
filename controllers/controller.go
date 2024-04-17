package controllers

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pradytpk/go-echo-mongodb/handlers"
	"github.com/pradytpk/go-echo-mongodb/middleware"
)

var e = echo.New()
var v = validator.New()

// Start starts the application
func Start() {
	//Env variable
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}
	e.Use(middleware.ServerMessage)
	e.GET("/", handlers.HealthCheck)
	e.GET("/products", handlers.GetProducts)
	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)

	e.Logger.Print("listening on port:", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
