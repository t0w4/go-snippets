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
	if err != nil {
		fmt.Fprintf(os.Stdout, "create file error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Fprintf(os.Stdout, "error writing record to csv: %v\n", err)
			os.Exit(1)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	os.Exit(0)
}
