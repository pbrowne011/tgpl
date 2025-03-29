// Exercise 1.4

package main

import (
	"bufio"
	"fmt"
	"os"
)

type LineInfo struct {
	Count     int
	FileNames []string
}

func main() {
	counts := make(map[string]*LineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, line_entry := range counts {
		n := line_entry.Count
		seen := make(map[string]int)
		fmt.Printf("%d\t%s\nFilenames:\n", n, line)
		for _, name := range line_entry.FileNames {
			if seen[name] == 0 {
				fmt.Printf("%s\n", name)
				seen[name] = 1
			}
		}
		fmt.Println()
	}
}

func countLines(f *os.File, counts map[string]*LineInfo, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		_, exists := counts[input.Text()]
		if !exists {
			counts[input.Text()] = &LineInfo{
				Count:     0,
				FileNames: make([]string, 0, 100),
			}
		}
		counts[input.Text()].Count++
		counts[input.Text()].FileNames = append(counts[input.Text()].FileNames, name)
	}
	// NOTE: ingoring potential errors from input.Err()
}
