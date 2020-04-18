package main

import (
	"io"
	"os"
)

func main()  {
	oldFile, oEerr := os.Open("old.txt")
	if oEerr != nil {
		panic(oEerr)
	}
	newFile, nErr := os.Create("new.txt")
	if nErr != nil {
		panic(nErr)
	}
	io.Copy(newFile, oldFile)
}
