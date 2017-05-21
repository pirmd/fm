package main

import (
	"fm/tui"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

var Cfg Config

type Config struct {
	// General
	Log *log.Logger // Logger of informative message for user
	// Colorscheme
	StyleNormal    tui.Style
	StyleError     tui.Style
	StyleHighlight tui.Style
	// Keybindings
	Keys map[string]func()
}

func init() {
	Cfg = Config{
		Log:            log.New(os.Stderr, "", log.Ltime),
		StyleNormal:    tui.NewStyle(),
		StyleError:     Cfg.StyleNormal.Fg(termbox.ColorRed),
		StyleHighlight: Cfg.StyleNormal.Bg(termbox.ColorBlack),
		Keys: map[string]func(){
			"j": func() { FM.NextWorkingFile() },
			"k": func() { FM.PrevWorkingFile() },
			"l": func() { FM.OpenWorkingFile() },
			"h": func() { FM.CdUp() },
			"q": func() { termbox.Interrupt() },
		},
	}
}
