package main

import (
	"flag"
	"fmt"
	"github.com/nitishfy/wc-tool/internals"
)

func main() {
	countBytes := flag.Bool("c",false,"count the number of bytes in a file")
	countLines := flag.Bool("l",false, "count the number of lines in a file")
	countWords := flag.Bool("w", false, "count the number of words in a file")
	countChars := flag.Bool("m", false, "count the number of characters in a file")
	flag.Parse()

	filepaths := flag.Args()
	
	for _, file := range filepaths {
		exists := internals.FileExists(file)
		if !exists {
			fmt.Printf("wc: %s No such file or directory\n", file)
			break
		}

		bytes := internals.CountBytes(file)
		lines, words, char := internals.Count(file)

		// count the number of bytes
		if *countBytes {
			fmt.Printf("%d %s\n",bytes, file)
		}

		// count the number of lines/words
		if *countLines || *countWords  || *countChars{
			if *countLines {
				fmt.Printf("%d %s\n", lines, file)
			}

			if *countWords {
				fmt.Printf("%d %s\n", words, file)
			}
			
			if *countChars {
				fmt.Printf("%d %s\n", char, file)
			}
		}
	}
}

