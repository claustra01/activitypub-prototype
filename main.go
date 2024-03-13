package main

import (
	"log"
	"os"

	"github.com/claustra01/activitypub-prototype/nodeinfo"
	"github.com/claustra01/activitypub-prototype/user"
	"github.com/claustra01/activitypub-prototype/wellknown"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/.well-known/nodeinfo", wellknown.NodeInfo)
	e.GET("/.well-known/webfinger", wellknown.WebFinger)

	e.GET("/nodeinfo/2.0", nodeinfo.NodeInfo2_0)
	e.GET("/nodeinfo/2.1", nodeinfo.NodeInfo2_1)

	e.GET("/users/test", user.MockUser)

	// log.Fatal(e.Start(":" + os.Getenv("PORT")))
	log.Fatal(e.StartTLS(":"+os.Getenv("PORT"), "server.crt", "server.key"))
}
