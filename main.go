package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func main() {
	a := app.New()
	w := a.NewWindow("UUID GEN")

	str := binding.NewString()
	str.Set("")

	text := widget.NewLabelWithData(str)

	w.Resize(fyne.NewSize(560, 480))
	button := widget.NewButton("New Uuid", func() {
		str.Set(uuid.New().String())
	})

	content := container.NewBorder(nil, nil, button, nil, text)
	w.SetContent(content)
	w.ShowAndRun()
}
