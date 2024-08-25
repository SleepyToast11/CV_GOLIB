package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: {input.pgm} {output.pgm} operation")
		os.Exit(1)
	}
	inputFile := args[0]
	outputFile := args[1]
	operation := args[2]

}
