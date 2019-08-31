package main

import (
	"github.com/skratchdot/open-golang/open"
)

// FileManager represents the file-manager application.
// It ensures the animation of the user interface while interacting with the file system.
// FileManager is also in charge to listen to the application events and to dispatch or manage them.
type FileManager struct {
	// Pointer to the user interface handler
	ui *FmTUI
	// Information about the foldes to be displayed
	wd   *Folder // Working dir
	upwd *Folder // Working dir parent (or nil)
}

// NewFileManager creates a new FileManager struct.
func NewFileManager(ui *FmTUI) *FileManager {
	return &FileManager{ui: ui}
}

// Open initiates a browsing session.
// Open takes an url as argument and recognize the protocol (scheme) supported by the vfs package
func (f *FileManager) Open(url string) {
	f.cd(NewFolderFromURL(url).SortByName())
}

// cd changes the browing current location to a new Folder.
func (f *FileManager) cd(wd *Folder) {
	if f.wd != nil {
		wd.SetWf(f.wd.Pwd())
	}
	f.wd = wd.SortByName()
	f.upwd = f.wd.Parent().SortByName()
	f.upwd.SetWf(f.wd.Pwd())
	f.showWd()
}

// CdUp offers a convenient shortcut to change ot the parent folder
func (f *FileManager) CdUp() {
	f.cd(f.upwd)
}

// OpenWorkingFile opens the selected file or folder
func (f *FileManager) OpenWorkingFile() {
	wf := f.wd.Pwf()
	if wf.IsDir() {
		f.cd(wf)
	} else {
		f.osopen(wf.Pwd())
	}
}

// osopen opens a File using the OS default program relying
// on launchers like (xdg-open or Start.exe). Magic is done by go-open package
func (f *FileManager) osopen(path string) {
	if err := open.Run(path); err != nil {
		// BUG: log to Sdterr (default config) does not play well with temrbox gui
		// Cfg.Log.Println("Fail to open", path, ":", err)
		fprintErr(f.ui.Info, err)
	}
}

// NextWorkingFile offers a convenient shortcut to move to the next file in the current folder
func (f *FileManager) NextWorkingFile() {
	f.wd.NextWf()
	f.showWf()
}

// PrevWorkingFile offers a convenient shortcut to move to the previous file in the current folder
func (f *FileManager) PrevWorkingFile() {
	f.wd.PrevWf()
	f.showWf()
}

// setcurdir displays the list of files of the current folder
func (f *FileManager) showWd() {
	f.ui.Wd.Reset()
	fprintDirLst(f.ui.Wd, f.wd)
	f.showTitle()
	f.showWf()
	f.showUpwd()
}

// showTitle prints title to understand what we are browsing
func (f *FileManager) showTitle() {
	f.ui.Title.Reset()
	fprintDirNfo(f.ui.Title, f.wd)
}

// showWf prints information about the current file
func (f *FileManager) showWf() {
	f.ui.Wd.SetCursor(f.wd.Index())
	f.ui.Info.Reset()
	fprintFiNfo(f.ui.Info, f.wd.GetWf())
	f.showPreview()
}

// showPreview displays a preview of the currently highlighted file
func (f *FileManager) showPreview() {
	wf := f.wd.Pwf()
	f.ui.Preview.Reset()
	if wf.IsDir() {
		fprintDirLst(f.ui.Preview, wf.SortByName())
	} else {
		//TODO/BUG: will not work if wf belong to a real virtual FileSystem
		f.ui.Preview.Reset()
		fprintContent(f.ui.Preview, wf.Pwd())
	}
}

// showUpwd displays the list of files of the parent folder
func (f *FileManager) showUpwd() {
	f.ui.Upwd.Reset()
	fprintDirLst(f.ui.Upwd, f.upwd)
	f.ui.Upwd.SetCursor(f.upwd.Index())
}
