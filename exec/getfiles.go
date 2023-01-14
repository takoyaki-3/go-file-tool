package main

import (
	"log"

	csvtag "github.com/takoyaki-3/go-csv-tag/v3"
	filetool "github.com/takoyaki-3/go-file-tool/v2"
)

func main() {

	fileList, err := filetool.DirWalk("../", filetool.DirWalkOption{
		Deep: 2,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = csvtag.DumpToFile(fileList, "files.csv")
	if err != nil {
		log.Fatalln(err)
	}
}
