// Dup2 prints the count, text of lines and name of file that appear more than once
// It reads from a list of named files in the input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//counts := make(map[string]int)
	//var filesWithDup []string
	files := os.Args[1:]
	if len(files) == 0 {
		os.Exit(2)
	} else {
		//fmt.Println("Files with dups: ")
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			data := countLines(f)
			if containsDups(data) {
				displayDups(f.Name(), data)
			}
			f.Close()
		}
	}
}

func countLines(f *os.File) map[string]int {
	data := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		data[input.Text()]++
	}
	return data
	// NOTE: ignoring potential errors from input.Err()
}
func containsDups(data map[string]int) bool {
	for _, n := range data {
		if n > 1 {
			return true
		}
	}
	return false
}
func displayDups(f string, counts map[string]int) {
	fmt.Printf("Filename: %s\n", f)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
