package main

import (
	"archive/zip"
	"os"
)

func main() {

	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}

	z := zip.NewWriter(file)
	defer z.Close()

	// how to use strings.Reader?

}
