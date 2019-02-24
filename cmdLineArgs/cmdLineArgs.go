package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Executable=" + os.Args[0])

	// Echo command line args
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println()
	fmt.Println("Args passed in:")
	fmt.Println(s)

	// Print out each command line arg and its index
	fmt.Println()
	fmt.Println("Args and their indexes:")
	for i, arg := range os.Args[1:] {
		fmt.Printf("%s=%d\n", arg, i)
	}

	fmt.Println("Done")
}
