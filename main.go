package main

import (
	"log"
	"os"

	"github.com/go-sharp/cerberus/v2"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var version = "v0.0.0"

func Version() string {
	return version
}

func main() {

	// Check if we are called with command line arguments
	if len(os.Args) >= 2 {
		if os.Args[1] == "run" && len(os.Args) >= 3 {

			if err := cerberus.RunService(os.Args[2]); err != nil {
				log.Fatalln(err)
			}
			os.Exit(0)
		}

	} else {

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

		app.Bind(svcs)
		app.Bind(Version)
		app.Run()
	}
}
