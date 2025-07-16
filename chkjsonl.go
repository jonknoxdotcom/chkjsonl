package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		// Open file (if possible)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Unable to open file " + filename)
			continue
		}
		defer file.Close()

		fmt.Print("Processing " + filename + " ")

		// read each line and ensure it's a valid JSON document
		scanner := bufio.NewScanner(file)
		lines := 1
		for scanner.Scan() {
			if json.Valid([]byte(scanner.Text())) {
				fmt.Print(".")
			} else {
				fmt.Printf("(%d)", lines)
			}
			lines++
		}

		fmt.Println()
	}
}
