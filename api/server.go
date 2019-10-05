package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/takatori/youtube-manager/api/middlewares"
	"github.com/takatori/youtube-manager/api/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YouTubeService())
	// Routes
	routes.Init(e)
	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
