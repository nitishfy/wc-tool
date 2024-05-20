package main

import (
	"flag"
	"fmt"

	"github.com/nitishfy/wc-tool/files"
)

func main() {
	bytesCount := flag.Bool("c",false,"count the number of bytes")
	linesCount := flag.Bool("l", false, "count the number of lines")
	wordsCount := flag.Bool("w", false, "count the number of words")
	flag.Parse()
	filepaths := flag.Args()
	for _, file := range filepaths {
		// check if file exists
		exists := files.FileExists(file)
		if !exists {
			fmt.Printf("wc: %s: No such file or directory\n", file)
			break
		}

		if *bytesCount {
			bytes := files.CountBytes(file)	
			fmt.Printf("%d %s\n",bytes,file)	
		}
		
		if *linesCount {
			lines := files.CountLines(file)
			fmt.Printf("%d %s\n",lines, file)
		}

		if *wordsCount {
			words := files.CountWords(file)
			fmt.Printf("%d %s\n", words, file)
		}

	}
}

