package menu

import ui "github.com/gizak/termui"

type Menu struct {
	HelpText  ui.Bufferer
	Shortcuts *Shortcuts
}

func NewMenu() *Menu {
	p := ui.NewPar(" PRESS q TO QUIT")
	p.Height = 2
	p.Width = 19
	p.TextFgColor = ui.ColorGreen
	p.Border = false
	return &Menu{
		HelpText:  p,
		Shortcuts: NewShortcuts(),
	}
}

func (m *Menu) Render() {
	ui.Render(m.HelpText)
	ui.Render(m.Shortcuts)
}

type Shortcuts struct {
	List *ui.List
}

func NewShortcuts() *Shortcuts {
	list := ui.NewList()
	list.Items = []string{
		"[c] containers",
		"[w] workers",
		"[v] volumes",
	}
	list.PaddingLeft = 3
	list.Border = true
	list.Height = 9
	list.Width = 20
	list.Y = 1
	list.BorderFg = ui.ColorWhite
	list.ItemFgColor = ui.ColorYellow
	return &Shortcuts{List: list}
}

func (s *Shortcuts) Highlight() {
	s.List.BorderFg = ui.ColorRed
}

func (s *Shortcuts) UnHighlight() {
	s.List.BorderFg = ui.ColorWhite
}

func (s *Shortcuts) Buffer() ui.Buffer {
	return s.List.Buffer()
}
