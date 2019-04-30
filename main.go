package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const pdfExt = ".pdf"

func main() {
	path := "."
	if len(os.Args) != 2 {
		fmt.Println("./pdf2text {path}")
	}
	path = os.Args[1]

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == pdfExt {
			fmt.Println(path)
		}

		return nil
	})

	panicOnErr(err)
}
