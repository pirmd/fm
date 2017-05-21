package main

import (
	"fm/vfs"
	"os"
	"path/filepath"
	"sort"
)


// Folder provides help to manipulate folders
type Folder struct {
	Ls    []os.FileInfo
	path  string
	fs    *vfs.FileSystem
	err   error
	index int
}

// NewFolderFromURL returns a new Folder for the given url 
func NewFolderFromURL(rawURL string) *Folder {
	fs, path, err := vfs.New(rawURL)
	if err != nil {
        return &Folder{err: err}
	}
	return NewFolder(fs, path)
}

// NewFolder returns a new Folder for the given path from the given file system
func NewFolder(fs *vfs.FileSystem, path string) *Folder {
	ls, err := fs.Readdir(path)
	return &Folder{Ls: ls, path: path, fs: fs, err: err}
}

// Parent generates a Folder instance for the parent folder
// Returns nil if no parent is found
func (f *Folder) Parent() *Folder {
	return NewFolder(f.fs, filepath.Dir(f.path))
}

// Child generates a Folder instance for one of the folder child
// Returns nil if no Child is found (index out of range)
func (f *Folder) Child(i int) *Folder {
	if (i < 0) || (i >= len(f.Ls)) {
		return nil
	}
	return NewFolder(f.fs, filepath.Join(f.path, f.Ls[i].Name()))
}

// Get returns a Folder element for the given index
// Returns nil if nothing is found (index out of range)
func (f *Folder) Get(i int) os.FileInfo {
	if (i < 0) || (i >= len(f.Ls)) {
		return nil
	}
	return f.Ls[i]
}

// Search looks for a given path within the Folder and set its working file to it if found
func (f *Folder) SetWf(path string) {
	basename := filepath.Base(path)
	for i, fi := range f.Ls {
		if fi.Name() == basename {
			f.index = i
			return
		}
	}
	f.index = 0
}

// SortByName sorts a filelist first by folder then by name
// TODO: does it worked if f.Ls is empty or if f is a nonexisting folder? 
func (f *Folder) SortByName() *Folder {
	if f != nil {
		sort.Slice(f.Ls, func(i, j int) bool {
			if f.Ls[i].IsDir() == f.Ls[j].IsDir() {
				return f.Ls[i].Name() < f.Ls[j].Name()
			}
			return f.Ls[i].IsDir()
		})
	}
	return f
}

// Pwd returns the Folder path
func (f *Folder) Pwd() string {
	return f.path
}

// Pwf returns the working file from the Folder (or nil)
func (f *Folder) Pwf() *Folder {
	return f.Child(f.index)
}

// GetWf returns information about the working file from the Folder (or nil)
func (f *Folder) GetWf() os.FileInfo {
	return f.Get(f.index)
}

// NextWf move the Folder working file to the next file
func (f *Folder) NextWf() {
	if f.index < len(f.Ls)-1 {
		f.index += 1
	}
}

// PrevWf move the Folder working file to the previous file
func (f *Folder) PrevWf() {
	if f.index > 0 {
		f.index -= 1
	}
}

// INdex return the current index value
func (f *Folder) Index() int {
	return f.index
}

// Error returns the current error status
func (f *Folder) Error() error {
	return f.err
}

// IsDir returns True if FOlder is really a dir
func (f *Folder) IsDir() bool {
	return f.Ls != nil
}
