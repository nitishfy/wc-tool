package internals

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"unicode/utf8"
)

// FileExists check if the file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func CountBytes(path string) int64 {
	file, _ := os.Stat(path)
	return file.Size()	 
}

func Count(path string) (int64,int64,int64) {
	file, err := OpenFile(path)
	if err != nil {
		fmt.Printf("error opening the file: %v", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	wordCount := 0
	charCount := 0

	for scanner.Scan() {
		// calculate number of lines
		lines++

		// calculate number of characters
		line := scanner.Text()
		charCount += utf8.RuneCountInString(line) + 1 // +1 to include the newline character

		// calculate number of words
		words := strings.Fields(line)
		wordCount += len(words) 
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	return int64(lines), int64(wordCount), int64(charCount)
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return &os.File{},nil
	}

	return file, err
}
