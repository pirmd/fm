package main

import (
	"fm/tui"
	"os"
    "github.com/dustin/go-humanize"
)

// FmTUI implements a tui.TerminalUI adapted to browse a filesystem
// FmTUI does not support browsing logics or commands, it just a
// container to provide the main ui parts
type FmTUI struct {
	*tui.TerminalUI
	Title   *tui.TextWidget
	Upwd    *tui.ListWidget
	Wd      *tui.ListWidget
	Preview *tui.TextWidget
	Info    *tui.TextWidget
}

// NewFmTUI returns a new FmTUI to play with
func NewFmTUI() (*FmTUI, error) {
	ui, err := tui.New()
	if err != nil {
		return nil, err
	}
	t := &FmTUI{
		TerminalUI: ui,
		Title:      tui.NewTextWidget(),
		Upwd:       tui.NewListWidget(),
		Wd:         tui.NewListWidget(),
		Preview:    tui.NewTextWidget(),
		Info:       tui.NewTextWidget(),
	}
	t.layout()
	return t, nil
}

// layout defines the layout principles for our file-manager application
func (t *FmTUI) layout() {
	t.AddWidget(t.Title)
	t.Upwd.SetStyle(Cfg.StyleNormal, Cfg.StyleHighlight)
	t.AddWidget(t.Upwd)
	t.Wd.SetStyle(Cfg.StyleNormal, Cfg.StyleHighlight)
	t.AddWidget(t.Wd)
	t.Preview.SetStyle(Cfg.StyleNormal, Cfg.StyleNormal)
	t.AddWidget(t.Preview)
	t.Info.SetStyle(Cfg.StyleNormal, Cfg.StyleNormal)
	t.AddWidget(t.Info)
	t.Resize()
}

// Resize set the size of the different ui Widgets depending on the main window size
func (t *FmTUI) Resize() {
	t.TerminalUI.Resize()
	if t.W < 2 || t.H < 2 {
		Cfg.Log.Fatal("Terminal is too small to do anything useful")
	}
	c := int(float32(t.W-2) / 4.0)
	t.Title.Resize(0, 0, t.W, 1)
	t.Upwd.Resize(0, 1, c, t.H-2)
	t.Wd.Resize(c+1, 1, 2*c, t.H-2)
	t.Preview.Resize(3*c+2, 1, t.W-3*c-2, t.H-2)
	t.Info.Resize(0, t.H-1, t.W, 1)
}

// fprintDirNfo formats folder information
func fprintDirNfo(w tui.ColorWriter, dir *Folder) {
	Cfg.StyleNormal.Fprintln(w, dir.Pwd())
}

// fprintLs displays a list of files in a widget
func fprintDirLst(w tui.ColorWriter, dir *Folder) {
	if dir.Error() != nil {
		fprintErr(w, dir.Error())
		return
	}
	if len(dir.Ls) == 0 {
		Cfg.StyleNormal.Fprintln(w, "<empty>")
	} else {
		for _, fi := range dir.Ls {
			fprintFiName(w, fi)
		}
	}
}

// fprintFiName formats file name according to its type
func fprintFiName(w tui.ColorWriter, fi os.FileInfo) {
	if fi.IsDir() {
		Cfg.StyleNormal.Bold().Fprintln(w, fi.Name()+"/")
	} else {
		Cfg.StyleNormal.Fprintln(w, fi.Name())
	}
}

// fprintFiNfo formats file information
func fprintFiNfo(w tui.ColorWriter, fi os.FileInfo) {
	if fi != nil {
		Cfg.StyleNormal.Fprintln(w, fi.Mode(), humanize.Bytes(uint64(fi.Size())), humanize.Time(fi.ModTime()))
	}
}

// fprintContent prints the content of the file identified by path
func fprintContent(w tui.ColorWriter, path string) {
	Cfg.StyleNormal.Fprintln(w, "<preview not implemented>")
	//TODO: Might use io.LimitReader
}

// fprintErr formats an error
func fprintErr(w tui.ColorWriter, err error) {
	Cfg.StyleError.Fprintln(w, err)
}
