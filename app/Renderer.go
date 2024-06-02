package app

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
  templates *template.Template
}

func (self *Renderer) New(assets *embed.FS) *Renderer {
  self.templates = template.Must(template.ParseFS(assets, "assets/*.html"))
  return self
}

func (self *Renderer) Render(wr io.Writer, name string, data interface{}, c echo.Context) error {
  return self.templates.ExecuteTemplate(wr, name, data)
}
