package tui

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// Style is a struct that facilitates use of colors and style for terminal printing
type Style struct {
	fg, bg termbox.Attribute
}

// NewStyle createsa  new Style with default colors
func NewStyle() Style {
	return Style{fg: termbox.ColorDefault, bg: termbox.ColorDefault}
}

// Fg set foreground color
func (s Style) Fg(fg termbox.Attribute) Style {
	s.fg = fg
	return s
}

// Bg set Background color
func (s Style) Bg(bg termbox.Attribute) Style {
	s.bg = bg
	return s
}

// Bold put foreground in Bold
func (s Style) Bold() Style {
	s.fg |= termbox.AttrBold
	return s
}

// Reverse reverse colors
func (s Style) Reverse() Style {
	s.fg |= termbox.AttrReverse
	return s
}

// Fprintln is a convenient wrapper that print in color to a window
// that implements ColorWriter interface
func (s Style) Fprintln(w ColorWriter, a ...interface{}) (n int, err error) {
	w.SetWriterColors(s.fg, s.bg)
	return fmt.Fprintln(w, a...)
}

// Fprint is a convenient wrapper that prints in color to a window
// that implements ColorWriter interface
func (s Style) Fprint(w ColorWriter, a ...interface{}) (n int, err error) {
	w.SetWriterColors(s.fg, s.bg)
	return fmt.Fprint(w, a...)
}
