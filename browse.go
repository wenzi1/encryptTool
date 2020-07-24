package main

import (
	"github.com/lxn/walk"
	"log"
	"os"
	"path/filepath"
)

type browseMw struct {
	*walk.MainWindow
	edit *walk.LineEdit
	path string
}

func (this *browseMw) getpath() {

	dlg := new(walk.FileDialog)

	absPath, _ := filepath.Abs(os.Args[0])
	dlg.FilePath = absPath
	dlg.Title = "Select File"
	dlg.Filter = "All files (*.*)|*.*"

	if ok, err := dlg.ShowBrowseFolder(this); err != nil {
		log.Fatalln("Error : File Open\n", err)
		return
	} else if !ok {
		return
	}
	this.path = dlg.FilePath
	path.SetText(this.path)
}
