package main 

//forloop

func Sum(numbers [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}




//use range , it returns two values, index and values 

func SumExample(numbers []int) int {
	sum := 0

	for _, number := range numbers{
		sum += number
	}
	return sum
}


func SumAll(numbersToSum ...[]int) []int{
    var sums []int

   for  _, numbers := range numbersToSum{
	   sums = append(sums, SumExample(numbers))
   }
   return sums
}


func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum{
		if len(numbers) == 0{
			sums = append(sums, 0)
		}else{
          tail := numbers[1:]
		sums = append(sums, SumExample(tail))
		}
		
	}
	return sums
}