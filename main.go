package main

import (
	ui "github.com/gizak/termui"
	"github.com/topherbullock/flight-instruments/menu"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	mainMenu := menu.NewMenu()
	mainMenu.Render()

	display := ui.NewPar("")
	display.Float = ui.AlignTop
	display.BorderFg = ui.ColorGreen
	display.TextFgColor = ui.ColorBlue
	display.Width = 20
	display.Height = 3
	display.X = 20
	ui.Render(display)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/v", func(ui.Event) {
		ui.Clear()
		mainMenu.Render()
		display.Text = "Volumes!"
		ui.Render(display)
	})

	ui.Handle("/sys/kbd/c", func(ui.Event) {
		ui.Clear()
		mainMenu.Render()
		display.Text = "Containers!"
		ui.Render(display)
	})
	ui.Handle("/sys/kbd/w", func(ui.Event) {
		ui.Clear()
		mainMenu.Render()
		display.Text = "Workers!"
		ui.Render(display)
	})

	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		ui.Clear()
		mainMenu.Render()
		ui.Render(display)
	})

	ui.Loop()
}
