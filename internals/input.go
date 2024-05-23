package internals

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"fmt"
)

type Handler struct {
	Reader io.Reader
	File string
}

type Options struct {
	CountBytes bool
	CountLines bool
	CountChars bool
	CountWords bool
}

func NewHandler(reader io.Reader, file string) *Handler {
	return &Handler{
		Reader: reader,
		File: file,
	}
}

func ProcessFiles(files []string, options Options) {
	for _, file := range files {
		fi, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()
		handler := NewHandler(fi, file)
		handler.Process(options)
	}
}

func (h *Handler) Process(options Options) {
	var bytes, lines, words, chars int
	scanner := bufio.NewScanner(h.Reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		bytes += len(line) + 1
		chars += len([]rune(line)) + 1 

		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			words++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	h.PrintCounts(bytes,lines,chars,words, options)	
}

func (h *Handler) PrintCounts(bytes,lines,chars,words int, options Options) {
	if h.File != "" {
		fmt.Printf("=======>%s\n", h.File)
	}
	
	if options.CountBytes {
		fmt.Printf("Bytes: %d\n", bytes)
	}

	if options.CountLines {
		fmt.Printf("Lines: %d\n", lines)
	}

	if options.CountChars {
		fmt.Printf("Chars: %d\n", chars)
	}

	if options.CountWords {
		fmt.Printf("Words: %d\n", words)
	}

	if !options.CountLines && !options.CountWords && !options.CountBytes && !options.CountChars {
		fmt.Printf("Lines:%d\nWords:%d\nBytes:%dChars:%d\n", lines, words, bytes, chars)
	}

	fmt.Println()
}