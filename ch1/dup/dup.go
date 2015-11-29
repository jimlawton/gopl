package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	cfiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, cfiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) == 0 {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				fmt.Printf("%d\t%s\t%v\n", n, line, cfiles[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		seen := false
		for _, fname := range files[input.Text()] {
			if fname == f.Name() {
				seen = true
				break
			}
		}
		if !seen {
			files[input.Text()] = append(files[input.Text()], f.Name())
		}
	}
}
