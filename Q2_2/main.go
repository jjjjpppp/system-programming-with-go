package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	csvData := [][]string{
		{"1","2","3"},
		{"4","5","6"},
	}
	writer.WriteAll(csvData)
	writer.Flush()
}
