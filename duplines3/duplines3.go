package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countLines(data []byte, counts map[string]int) {
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}
}

func printDupLines(counts map[string]int) {
	// Print out lines that are duplicates
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplines3: %v\n", err)
		}
		countLines(data, counts)
	}

	printDupLines(counts)
}
