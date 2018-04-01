package main

import (
	"fmt"
	"bufio"
	"os"
	"bytes"
	"strconv"
)

func main() {
	var output bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	_, _ = reader.ReadString('\n')
	var input, _ = reader.ReadString('\n')

	var rotateStr, _ = reader.ReadString('\n')
	var rotate, _ = strconv.Atoi(rotateStr)

	for _, ch := range input {
		if isLetter(ch) {
			var lowerLimit int

			if ch >= 'a' {
				lowerLimit = int('a')
			} else {
				lowerLimit = int('A')
			}

			encryptedStr := string(lowerLimit + (int(ch)+rotate-lowerLimit)%26)
			output.WriteString(encryptedStr)
		} else {
			output.WriteString(string(ch))
		}
	}

	fmt.Println(output.String())
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z');
}
