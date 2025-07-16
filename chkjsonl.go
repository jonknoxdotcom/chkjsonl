package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	for _, filename := range argsWithoutProg {
		// Open file (if possible)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Unable to locate file " + filename)
			continue
		}
		defer file.Close()

		// read each line and validate
		fmt.Print("Processing " + filename + " ")

		scanner := bufio.NewScanner(file)
		lines := 1
		for scanner.Scan() {
			if !json.Valid([]byte(scanner.Text())) {
				fmt.Printf("(%d)", lines)
			} else {
				fmt.Print(".")
			}
			lines++
		}

		fmt.Println("")
	}
}
