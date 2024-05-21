package main

import (
	"flag"
	"fmt"
)

func main() {
	countBytes := flag.Bool("c", false,"count the number of bytes in a file")
	countLines := flag.Bool("l", false, "count the number of lines in a file")
	countWords := flag.Bool("w", false, "count the number of words in a file")
	countChars := flag.Bool("m", false, "count the number of characters in a file")	
	flag.Parse()

	// default case when no argument is passed
	if !*countBytes || !*countLines || !*countWords {
		*countBytes, *countLines, *countChars = true, true, true
	}
	file := flag.Arg(0)
	if file == "" {
		fmt.Printf("Please provide the file path")
		return
	}

}

