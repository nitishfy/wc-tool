package interfaces

type WC interface {
	CountBytes() int
	CountLines() int
	CountWords() int
	CountChars() int
}