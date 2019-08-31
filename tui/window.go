package tui

import (
	"github.com/nsf/termbox-go"
)

// Window represents a window of the user interface
type Window struct {
	// position of top-left window corner
	x0, y0 int
	// size of window
	w, h int
	// Window normal and highlight colorscheme
	fg, bg   termbox.Attribute
	hfg, hbg termbox.Attribute
}

// NewWindow creates new Window
func NewWindow() *Window {
	return &Window{}
}

// Resize modify the size of the window
func (win *Window) Resize(x0, y0, w, h int) {
	win.x0, win.y0, win.w, win.h = x0, y0, w, h
}

// SetLine prints a line in color using the normal colorscheme
func (win *Window) SetLine(y int, line []termbox.Cell) {
	for x, cell := range line {
		if x == win.w {
			termbox.SetCell(win.x0+x-1, win.y0+y, '…', cell.Fg, cell.Bg)
			return
		}
		termbox.SetCell(win.x0+x, win.y0+y, cell.Ch, cell.Fg, cell.Bg)
	}
}

// SetLineHighlight prints a line in color using the highlight colorscheme
func (win *Window) SetLineHighlight(y int, line []termbox.Cell) {
	for x, cell := range line {
		if x == win.w {
			termbox.SetCell(win.x0+x-1, win.y0+y, '…', cell.Fg, win.hbg)
			return
		}
		termbox.SetCell(win.x0+x, win.y0+y, cell.Ch, cell.Fg, win.hbg)
	}
	for x := len(line); x < win.w; x++ {
		termbox.SetCell(win.x0+x, win.y0+y, ' ', win.hfg, win.hbg)
	}
}

// SetStyle defines the window normal and highlight colorscheme
func (win *Window) SetStyle(nStyle, hStyle Style) {
	win.fg, win.bg = nStyle.fg, nStyle.bg
	win.hfg, win.hbg = hStyle.fg, hStyle.bg
}

// Clear clears the window content
func (win *Window) Clear() {
	for y := 0; y < win.h; y++ {
		for x := 0; x < win.w; x++ {
			termbox.SetCell(win.x0+x, win.y0+y, ' ', win.fg, win.bg)
		}
	}
}
