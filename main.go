package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nitishfy/wc-tool/file_handler"
	"github.com/nitishfy/wc-tool/interfaces"
	"github.com/nitishfy/wc-tool/stdin_handler"
)

func main() {
	countBytes := flag.Bool("c", false,"count the number of bytes in a file")
	countLines := flag.Bool("l", false, "count the number of lines in a file")
	countWords := flag.Bool("w", false, "count the number of words in a file")
	countChars := flag.Bool("m", false, "count the number of characters in a file")	
	flag.Parse()

	file := flag.Arg(0)
	var handler interfaces.WC

	if file != "" {
		handler = &file_handler.FileHandler{
			FilePath: file,
		}

		// check if file exists locally
		exists := file_handler.FileExists(file)
		if !exists {
			fmt.Printf("%s does not exist\n", file)
			return 
		}
	} else {
		// check from stdin
		isEmpty, err := EmptyStdin()
		if err != nil {
			log.Fatal(err)
		}

		// stdin is empty too - it means no file has been passed
		if isEmpty {
			fmt.Printf("Pass a file as an argument")
			return 
		}
		
		handler = &stdin_handler.StdinHandler{}
	}

	switch {
	case *countBytes:
		fmt.Printf("Bytes: %d\n", handler.CountBytes())
	case *countLines:
		fmt.Printf("Lines: %d\n", handler.CountLines())
	case *countChars:
		fmt.Printf("Chars: %d\n", handler.CountChars())	
	case *countWords:
		fmt.Printf("Words: %d\n", handler.CountWords())
	default:
		fmt.Printf("%d %d %d ", handler.CountLines(), handler.CountWords(), handler.CountBytes())			
	}

	fmt.Printf("%s", file)
}

// EmptyStdin checks if the Stdin is empty
func EmptyStdin() (bool, error) {
  fi, err := os.Stdin.Stat()
  if err != nil {
    log.Fatal(err)
  }
  
  if fi.Mode() & os.ModeNamedPipe == 0 {
    return true, nil
  } 

  return false, nil
}
