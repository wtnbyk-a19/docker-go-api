package main

import (
	"docker-go-api/src/infrastructure/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load("/go/api/.env")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Routes
	router.RouterInit(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
