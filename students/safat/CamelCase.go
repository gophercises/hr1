package main

import "fmt"

func main() {
	var line string

	fmt.Scanln(&line)

	wordCount := 0

	if len(line) > 0 {
		wordCount = 1
	}

	for _, ch := range line {
		if ch >= 'A' && ch <= 'Z' {
			wordCount++
		}
	}

	fmt.Println(wordCount)
}
