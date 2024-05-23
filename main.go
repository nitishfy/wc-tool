package main

import (
	"flag"
	"os"
	"github.com/nitishfy/wc-tool/internals"
)

func main() {
	countBytes := flag.Bool("c", false, "count the number of bytes")
	countLines := flag.Bool("l", false, "count the number of lines")
	countWords := flag.Bool("w", false, "count the number of words")
	countChars := flag.Bool("m", false, "count the number of chars")
	flag.Parse()

	options := internals.Options{
		CountBytes: *countBytes,
		CountLines: *countLines,
		CountChars: *countChars,
		CountWords: *countWords,
	}

	files := flag.Args()

	if len(files) == 0 {
		handler := internals.NewHandler(os.Stdin,"")
		handler.Process(options)
	} else {
		internals.ProcessFiles(files,options)
	}
}