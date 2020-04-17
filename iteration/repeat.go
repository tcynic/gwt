package iteration

// Repeat takes a string and returns that string 5 times
func Repeat (character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}