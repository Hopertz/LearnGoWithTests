package Iteration

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated = repeated + character
	}

	return repeated

}
