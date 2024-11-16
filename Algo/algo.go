package main

import (
	"fmt"
)


func BinarySearch(numberArr []int, number int ) bool{
     size := len(numberArr)

     if size <= 0 {
      return false
     }

     low := 0
     high := size - 1

     for low <= high {
      
          mid := low + (high - low ) / 2 
           fmt.Printf("low: %d, mid: %d, high: %d\n", low, mid, high) // Debug log

          if(numberArr[mid] == number){
            return true
          }

             if number < numberArr[mid]{
              high = mid - 1
             }else
             {
              low = mid + 1 
             }
     }
  return false
}





func InsertionSort(numberToSort ...[]int) []int{

  return []int{1,2,3,4,5,6}

}