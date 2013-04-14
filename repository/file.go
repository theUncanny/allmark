// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"github.com/andreaskoch/allmark/path"
	"github.com/andreaskoch/allmark/watcher"
	"path/filepath"
)

type File struct {
	watcher.ChangeHandler
	path string
}

func NewFile(filePath string) (*File, error) {

	// create a file change handler
	fileChangeHandler, err := watcher.NewFileChangeHandler(filePath)
	if err != nil {
		return nil, fmt.Errorf("Could not create a change handler for file %q.\nError: %s\n", filePath, err)
	}

	// create the file
	file := &File{
		ChangeHandler: fileChangeHandler,
		path:          filePath,
	}

	return file, nil
}

func (file *File) String() string {
	return fmt.Sprintf("%s", file.path)
}

func (file *File) Path() string {
	return file.path
}

func (file *File) PathType() string {
	return path.PatherTypeFile
}

func (file *File) Directory() string {
	return filepath.Dir(file.Path())
}