package main

import (
	"embed"
	"gohtmx/app"
)

//go:embed assets/*
var assets embed.FS

func main() {
  app.Start(&assets)
}
