package main

type WC interface {
	CountBytes() 
	CountLines()
	CountWords() 
	CountChars()
}