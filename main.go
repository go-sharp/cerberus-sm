package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	svcs := &Services{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "cerberus-sm",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(basic)
	app.Bind(svcs)
	app.Run()
}
