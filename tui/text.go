package tui

import (
	"bytes"
	"github.com/nsf/termbox-go"
)

// TextWidget represents a scrollable window that displays colored text
// TextWidget implements the io.Writer interface to facilitate entering content
type TextWidget struct {
	// Text to be displayed supporting colors
	text [][]termbox.Cell
	// Position for viewing
	vx, vy int
	// Position for writing
	wx, wy int
	// Colors for writing
	wfg, wbg termbox.Attribute
	// TextWidget is a Window
	*Window
}

// NewTextWidget creates new TextWidget
func NewTextWidget() *TextWidget {
	return &TextWidget{Window: NewWindow()}
}

// Write corresponds to the implementataion of the io.Writer interface
func (txt *TextWidget) Write(p []byte) (n int, err error) {
	if txt.wy >= len(txt.text) {
		txt.text = append(txt.text, []termbox.Cell{})
	}
	for _, ch := range bytes.Runes(p) {
		switch ch {
		case '\n':
			txt.wy += 1
			txt.wx = 0
		default:
			if txt.wx == len(txt.text[txt.wy]) {
				txt.text[txt.wy] = append(txt.text[txt.wy], termbox.Cell{ch, txt.wfg, txt.wbg})
			} else {
				txt.text[txt.wy][txt.wx] = termbox.Cell{ch, txt.wfg, txt.wbg}
			}
			txt.wx += 1
		}
	}
	return len(p), nil
}

// Draw implements Widget interface to display the window content
func (txt *TextWidget) Draw() {
	var y, offsetX, offsetY int

	offsetX = txt.w * (txt.vx / txt.w)
	offsetY = txt.h * (txt.vy / txt.h)

	txt.Clear()

	for r, line := range txt.text {
		if (offsetX > len(line)-1) || (offsetY > r) {
			continue
		}
		y = r - offsetY
		if y > txt.h {
			break
		}
		if r == txt.vy {
			txt.SetLineHighlight(y, line[offsetX:])
		} else {
			txt.SetLine(y, line[offsetX:])
		}
	}
}

// Reset empties the text to be displayed
func (txt *TextWidget) Reset() {
	txt.text = [][]termbox.Cell{}
	txt.vx, txt.vy = 0, 0
	txt.wx, txt.wy = 0, 0
	txt.Clear()
}

// Scroll modify the viewing position of the text horizontally and/or vertically
func (txt *TextWidget) Scroll(x, y int) {
	//TODO: shall we also set/check vx limits?
	if x < 0 {
		return
	}
	if (y < 0) || (y > len(txt.text)-1) {
		return
	}
	txt.vx, txt.vy = x, y
}

// SetWriterColors modifies the colors for writing
func (txt *TextWidget) SetWriterColors(fg, bg termbox.Attribute) {
	txt.wfg, txt.wbg = fg, bg
}
