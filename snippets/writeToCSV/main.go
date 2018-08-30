package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("test.csv")
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
	}
	w := csv.NewWriter(file)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Fprintf(os.Stdout, "error writing record to csv: %v\n", err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}
