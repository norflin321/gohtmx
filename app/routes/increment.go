package routes

import "github.com/labstack/echo/v4"

func Increment(c echo.Context) error {
  data.Counter += 1
  return c.Render(200, "counter", data)
}
