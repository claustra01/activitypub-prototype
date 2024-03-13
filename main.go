package main

import (
	"io/ioutil"
	"log"

	"github.com/claustra01/activitypub-prototype/nodeinfo"
	"github.com/claustra01/activitypub-prototype/user"
	"github.com/claustra01/activitypub-prototype/wellknown"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func main() {

	var config Config
	yamlFile := "config.yml"

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("host", config.Host)
			return next(c)
		}
	})

	e.GET("/.well-known/nodeinfo", wellknown.NodeInfo)
	e.GET("/.well-known/webfinger", wellknown.WebFinger)

	e.GET("/nodeinfo/2.0", nodeinfo.NodeInfo2_0)
	e.GET("/nodeinfo/2.1", nodeinfo.NodeInfo2_1)

	e.GET("/users/test", user.MockUser)

	log.Fatal(e.Start(":" + config.Port))
	// log.Fatal(e.StartTLS(":"+config.Port, "server.crt", "server.key"))
}
