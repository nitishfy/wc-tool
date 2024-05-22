package stdin_handler

import (
	"bufio"
	"os"

)

type StdinHandler struct {}

// CountBytes counts the total number of bytes from Stdin
func (s *StdinHandler) CountBytes() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	var bytes int
	for scanner.Scan() {
		bytes++
	}

	return bytes
}

// CountLines counts the total number of lines from Stdin
func (s *StdinHandler) CountLines() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines int
	for scanner.Scan() {
		lines++
	}

	return lines
}

// CountWords counts the total number of words from stdin
func (s *StdinHandler) CountWords() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var words int
	for scanner.Scan() {
		words++
	}

	return words
}

// CountChars counts the total number of characters from stdin
func (s *StdinHandler) CountChars() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	var chars int
	for scanner.Scan() {
		chars++
	}

	return chars
}
