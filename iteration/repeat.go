package iteration

func Repeat(character string, amountToRepeat int) string {
	var repeated string
	for i := 0; i < amountToRepeat; i++ {
		repeated += character
	}
	return repeated
}
