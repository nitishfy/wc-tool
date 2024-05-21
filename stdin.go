package main

import (
	"bufio"
	"os"

)

type StdinHandler struct {
	count int
}

func (s *StdinHandler) CountBytes() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		s.count++
	}
}

func (s *StdinHandler) CountLines() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s.count++
	}
}

func (s *StdinHandler) CountWords() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s.count++
	}
}

func (s *StdinHandler) CountChars() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		s.count++
	}
}