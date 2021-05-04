package iteration

const repeatedCount = 5


func Repeat(character string, numberOfTimes int) string {
     var repeated string
	 for i := 0; i < numberOfTimes; i++ {
		 repeated += character
	 }
	return repeated
}