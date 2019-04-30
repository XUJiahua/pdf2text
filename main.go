package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const pdfExt = ".pdf"
const txtExt = ".txt"

func main() {
	path := "."
	if len(os.Args) != 3 {
		fmt.Println("./pdf2text {srcDir} {dstDir}")
		return
	}
	path = os.Args[1]
	dstDir := os.Args[2]

	err := os.MkdirAll(dstDir, os.ModePerm)
	panicOnErr(err)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == pdfExt {
			fmt.Printf("found %s\n", path)
		}

		_, file := filepath.Split(path)
		dstPath := filepath.Join(dstDir, strings.Trim(file, pdfExt)+txtExt)

		out, _ := exec.Command("pdftotext", path, dstPath).CombinedOutput()
		fmt.Printf("%s\n", out)
		return nil
	})

	panicOnErr(err)
}
