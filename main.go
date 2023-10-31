package main

import (
	"parking/scenes"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Begining")

	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1000, 800))
	scenes.NewParkingScene(w)
	w.ShowAndRun()
}
