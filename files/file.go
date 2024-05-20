package files

import (
	"bufio"
	"os"
	"fmt"
)

// fileExists check if a file exists or not
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func CountBytes(path string) int {
	file, _ := os.Stat(path)
	return int(file.Size())
}

func CountLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening the file: %v\n", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}
	
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lineCount
}	

func CountWords(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening the file: %v\n", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0

	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}