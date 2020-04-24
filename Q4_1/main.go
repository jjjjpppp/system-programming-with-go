package main

import (
	"fmt"
	"time"
)

func main() {
	seconds := 5
	ch := time.After(time.Duration(seconds) * time.Second)

	<-ch
	fmt.Printf("Waited %d Seconds", seconds)
}
