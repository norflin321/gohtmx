package routes

import "github.com/labstack/echo/v4"

type TData struct {
  Counter int
}

var data = TData{
  Counter: 0,
}

func Home(c echo.Context) error {
  data.Counter = 0
  return c.Render(200, "index", data)
}
