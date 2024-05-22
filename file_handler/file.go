package file_handler

import (
	"bufio"
	"os"

)

type FileHandler struct {
	FilePath string
	// count int 
}

// FileExists check if the file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

// OpenFile opens a file 
func (f *FileHandler) OpenFile() (*os.File, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, err
	}
	
	return file,nil
}

// CountBytes counts the total number of bytes in a file
func (f *FileHandler) CountBytes() int {
	file, err := f.OpenFile()
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	var bytes int
	for scanner.Scan() {
		bytes++
	}

	return bytes
}

// CountLines counts the total number of lines in a file
func (f *FileHandler) CountLines() int {
	file, err := f.OpenFile()
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines int
	for scanner.Scan() {
		lines++
	}

	return lines
}

// CountWords counts the total number of words in a file
func (f *FileHandler) CountWords() int {
	file, err := f.OpenFile()
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words int
	for scanner.Scan() {
		words++
	}

	return words
}

// CountChars counts the total number of characters in a file
func (f *FileHandler) CountChars() int {
	file, err := f.OpenFile()
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var chars int	
	for scanner.Scan() {
		chars++
	}

	return chars
}
