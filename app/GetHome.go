package app

import (
  "html/template"

  "github.com/labstack/echo/v4"
)

type THomeData struct {
  Counter int
  Css template.CSS
}

var HomeData = THomeData{
  Counter: 0,
}

func GetHome(c echo.Context) error {
  HomeData.Counter = 0
  return c.Render(200, "index", HomeData)
}
