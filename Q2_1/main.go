package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("format_test.txt")
	if err != nil {
		panic(err)
	}
	writer := io.Writer(file)
	fmt.Fprintf(writer,"%d\n",999)
	fmt.Fprintf(writer,"%s\n","Test write")
	fmt.Fprintf(writer,"%f\n",3.141519)
}