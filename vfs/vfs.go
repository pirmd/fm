package vfs

import (
	"errors"
	"github.com/spf13/afero"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Type FileSystem represents a FileSystem offering a standard set
// of functions to interact with
type FileSystem struct {
	afero.Fs
}

// New returns a new FileSystem whose type corresponds to the given rawURL
// as well as the path within this FileSystem pointed by the url
func New(rawURL string) (*FileSystem, string, error) {
	// Force default file:// scheme to avoid some misbehaviour of
	// url.Parse() when scheme is absent
	if strings.HasSuffix(rawURL, "//") {
		rawURL = "file:" + rawURL
	} else if !strings.ContainsAny(rawURL, "://") {
		rawURL = "file://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, rawURL, err
	}

	switch {
	case u.Scheme == "file":
		dir, err := filepath.Abs(u.Path)
		return &FileSystem{afero.NewOsFs()}, dir, err
	default:
		return nil, rawURL, errors.New("Unknown url scheme " + u.Scheme)
	}
}

// Readdir list the content of a directory
func (fs *FileSystem) Readdir(dirname string) ([]os.FileInfo, error) {
	return afero.ReadDir(fs, dirname)
}
