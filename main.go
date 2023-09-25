package main

import (
	_ "API-ECHO/docs"
	"API-ECHO/router"

	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title REST Echo Base
// @version 1.0
// @description This is an API base without database interaction

// @contact.name My linkedin profile
// @contact.url https://www.linkedin.com/in/joaoschoen/
// @contact.email joaoschoen@gmail.com

func main() {
	// Environment config
	//
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
		os.Exit(1)
	}
	// PORT
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		PORT = 3000
	}
	// DEBUG MODE
	DEBUG, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		DEBUG = false
	}

	// Echo instance
	Server := echo.New()

	// Routes
	router.InitRoutes(Server)

	// Print routes for debbuging
	if DEBUG {
		data, err := json.MarshalIndent(Server.Routes(), "", "  ")
		if err != nil {

		}
		os.WriteFile("routes.json", data, 0644)
		// Generate docs
		Server.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// Middleware stack
	Server.Use(middleware.CORS())
	Server.Use(middleware.Secure())
	Server.Use(middleware.RequestID())
	Server.Use(middleware.Logger())
	Server.Pre(middleware.RemoveTrailingSlash())
	Server.Use(middleware.Recover())

	// Initialize server
	Server.Logger.Fatal(Server.Start(fmt.Sprint(":", PORT)))
}
