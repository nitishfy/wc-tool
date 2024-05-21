package internals

import (
	"bufio"
	"os"

)

type FileHandler struct {
	FilePath string
	count int
}

// FileExists check if the file exists
func (f *FileHandler) FileExists() bool {
	if _, err := os.Stat(f.FilePath); err != nil {
		return false
	}

	return true
}

func (f *FileHandler) OpenFile() (*os.File, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, err
	}
	
	return file,nil
}

func (f *FileHandler) CountBytes() {
	file, err := f.OpenFile()
	if err != nil {
		return 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		f.count++
	}
}

func (f *FileHandler) CountLines() {
	file, err := f.OpenFile()
	if err != nil {
		return 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		f.count++
	}
}

func (f *FileHandler) CountWords() {
	file, err := f.OpenFile()
	if err != nil {
		return 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		f.count++
	}
}

func (f *FileHandler) CountChars() {
	file, err := f.OpenFile()
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		f.count++
	}
}
