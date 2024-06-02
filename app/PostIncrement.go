package app

import "github.com/labstack/echo/v4"

func PostIncrement(c echo.Context) error {
  HomeData.Counter += 1
  return c.Render(200, "counter", HomeData)
}
