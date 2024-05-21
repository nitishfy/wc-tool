package main

import (
	"flag"
	"fmt"
	"github.com/nitishfy/wc-tool/internals"
)

func main() {
	countBytes := flag.Bool("c", false,"count the number of bytes in a file")
	countLines := flag.Bool("l", false, "count the number of lines in a file")
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

		lines, words, char := internals.Count(file)	

		switch {
		case *countBytes:
			fmt.Printf("%d ", internals.CountBytes(file))
		case *countLines: 
			fmt.Printf("%d ", lines)
		case *countWords:
			fmt.Printf("%d ", words)
		case *countChars:
			fmt.Printf("%d ", char)
		default:
			fmt.Printf("%d %d %d ", lines, words, internals.CountBytes(file))				
		}

		fmt.Printf("%s\n", file)
	}
}

