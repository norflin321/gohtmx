package app

import (
	"embed"
	"html/template"

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


  css, _ := assets.ReadFile("assets/index.css")
  HomeData.Css = template.CSS(string(css))

  e.GET("/", GetHome)
  e.POST("/increment", PostIncrement)

  e.Logger.Fatal(e.Start(":1234"))
}
