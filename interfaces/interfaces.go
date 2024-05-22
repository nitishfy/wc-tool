package interfaces

// WC defines interface to calculate details from stdin and file path
type WC interface {
	CountBytes() int
	CountLines() int
	CountWords() int
	CountChars() int
}