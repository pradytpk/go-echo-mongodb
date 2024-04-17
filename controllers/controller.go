package controllers

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pradytpk/go-echo-mongodb/handlers"
	"github.com/pradytpk/go-echo-mongodb/middlewares"
)

func Start() {
	var e = echo.New()

	// Initialize and register validator
	e.Validator = &handlers.ProductValidator{Validator: validator.New()}

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
	e.Use(middlewares.ServerMessage)
	e.GET("/", handlers.HealthCheck)
	e.GET("/products", handlers.GetProducts)
	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)

	e.Logger.Print("listening on port:", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
