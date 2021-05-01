package integers

import "fmt"

//add takes two interfer and return the sum of them

func Add(x, y int) int {
	return x + y 
}

func ExampleAdd(){
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}