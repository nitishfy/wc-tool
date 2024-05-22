package stdin_handler

import (
	"bufio"
	"os"

)

type StdinHandler struct {
	// count int
}

func (s *StdinHandler) CountBytes() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	var bytes int
	for scanner.Scan() {
		bytes++
	}

	return bytes
}

func (s *StdinHandler) CountLines() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines int
	for scanner.Scan() {
		lines++
	}

	return lines
}

func (s *StdinHandler) CountWords() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var words int
	for scanner.Scan() {
		words++
	}

	return words
}

func (s *StdinHandler) CountChars() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	var chars int
	for scanner.Scan() {
		chars++
	}

	return chars
}
