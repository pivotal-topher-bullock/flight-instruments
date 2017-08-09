package main

import ui "github.com/gizak/termui"

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar("PRESS q TO QUIT")
	p.Height = 3
	p.Width = 15
	p.TextFgColor = ui.ColorGreen
	p.BorderFg = ui.ColorCyan

	ui.Render(p)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
