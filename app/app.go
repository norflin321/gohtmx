package app

import (
	"embed"
	"gohtmx/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(assets *embed.FS) {
  e := echo.New()
  e.HideBanner = true
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  renderer := &Renderer{}
  e.Renderer = renderer.New(assets)

  e.GET("/", routes.Home)
  e.POST("/increment", routes.Increment)

  e.Logger.Fatal(e.Start(":1234"))
}
