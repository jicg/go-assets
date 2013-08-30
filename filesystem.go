package assets

import (
	"bytes"
	"net/http"
	"os"
)

// An in-memory asset file system. The file system implements the
// http.FileSystem interface.
type FileSystem struct {
	// A map of directory paths to the files in those directories.
	Dirs map[string][]string

	// A map of file/directory paths to assets.File types.
	Files map[string]*File

	// Whether or not the file data in the file system is stored in gzip
	// compressed form.
	Compressed bool
}

// Implementation of http.FileSystem
func (f *FileSystem) Open(path string) (http.File, error) {
	if fi, ok := f.Files[path]; ok {
		if !fi.IsDir() {
			// Make a copy for reading
			ret := fi
			ret.buf = bytes.NewReader(ret.Data)

			return ret, nil
		}

		return fi, nil
	}

	return nil, os.ErrNotExist
}

func (f *FileSystem) readDir(path string) ([]os.FileInfo, error) {
	if d, ok := f.Dirs[path]; ok {
		ret := make([]os.FileInfo, len(d))

		for i, v := range d {
			ret[i] = f.Files[v]
		}

		return ret, nil
	}

	return nil, os.ErrNotExist
}