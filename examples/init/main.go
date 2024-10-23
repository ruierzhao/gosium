package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/ruierzhao/gosium"
)

func test_fyne() {
	app := app.NewWithID("gosium.example.init")
	window := app.NewWindow("gosium example")

	btn := widget.NewButton("click", func() {
		log.Println("Click")
		gosium.Viewer()
		if fulls := window.FullScreen(); fulls {
			window.SetFullScreen(false)
		} else {
			window.SetFullScreen(true)
		}
	})
	window.SetContent(btn)
	window.SetFullScreen(true)

	window.ShowAndRun()
	app.Quit()

}

func test_gosium() {
	gosium.Viewer()
}

func main() {
	test_gosium()
}
