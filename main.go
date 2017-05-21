package main

import (
	"github.com/nsf/termbox-go"
	"os"
)

var (
	FM *FileManager
)

func main() {
	UI, err := NewFmTUI()
	if err != nil {
		Cfg.Log.Fatal(err)
	}
	defer UI.Close()

	FM = NewFileManager(UI)
    if len(os.Args) == 1 {
        FM.Open(".")
    } else {
        FM.Open(os.Args[1])
	}

	if err := UI.Run(); err != nil {
		Cfg.Log.Fatal(err)
	}

	for {
		select {
		case evt := <-UI.UIEvents:
			switch evt.Type {
			case termbox.EventKey:
				if c, ok := Cfg.Keys[string(evt.Ch)]; ok {
					c()
				} else {
					Cfg.Log.Println("Key pressed", evt.Ch, "unknown")
				}
			case termbox.EventResize:
				UI.Resize()
			case termbox.EventInterrupt:
				return
			case termbox.EventError:
				Cfg.Log.Fatal(evt.Err)
			}
		}

		if err := UI.Draw(); err != nil {
			Cfg.Log.Fatal(err)
		}
	}
}
