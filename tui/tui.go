package tui 

import (
    "io"
    "github.com/nsf/termbox-go"
)

// Widget defines an interface for TerminalUI widgets
type Widget interface {
    // Draw widget content on screen
	Draw()
    // Set or reset Size of the Widget
    Resize(int, int, int, int)
}

// ColorWriter defines an interface allowing to define color 
// with an io.Writer interface
type ColorWriter interface {
    io.Writer
    // Define foreground and background color
    SetWriterColors(termbox.Attribute, termbox.Attribute)
}

type TerminalUI struct {
    W, H      int
    Widgets   []Widget
	UIEvents  chan termbox.Event
}

func New() (*TerminalUI, error) {
	if err := termbox.Init(); err != nil {
        return nil, err	
	}
    t := &TerminalUI{
        UIEvents: make(chan termbox.Event),
    }
    t.Resize()
    return t, nil
}

func (t *TerminalUI) Resize() {
    t.W, t.H = termbox.Size()
}

func (t *TerminalUI) Close() {
	termbox.Close()
}

func (t *TerminalUI) AddWidget(w Widget) {
    t.Widgets = append(t.Widgets, w)
}

func (t *TerminalUI) Draw() error {
	for _, w := range t.Widgets {
        w.Draw()
	}

	return termbox.Flush()
}

func (t *TerminalUI) Run() error {
	go func() {
		for {
			t.UIEvents <- termbox.PollEvent()
		}
	}()

    return t.Draw()
}
