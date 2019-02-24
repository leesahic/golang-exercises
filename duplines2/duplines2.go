package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int, filemap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		filemap[line] = append(filemap[line], f.Name())
		counts[line]++
	}
}

func main() {
	counts := make(map[string]int)
	filemap := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filemap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplines2: %v\n", err)
				continue
			}
			countLines(f, counts, filemap)
			f.Close()
		}
	}

	// Print out lines that are duplicates
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("\tFound in: %v", filemap[line])
		}
	}
}
