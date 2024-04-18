package iterations

func Repeat(character string, timesToRepeat int) string {
	var repeatedChar string
	for i := 0; i < timesToRepeat; i++ {
		repeatedChar += character
	}

	return repeatedChar
}
