package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {



	t.Run("Collection of 5 numbers", func(t *testing.T){
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15
	//%v for arrays

	if got != want {
		t.Errorf("got %d want %d  given, %v ", got, want, numbers)
	}})


	t.Run("collection of any size", func (t *testing.T)  {
		
        numbers := []int{1,2,3}

		got:= SumExample(numbers)
		want := 6

		if got != want{
			t.Errorf("got %d want %d  given, %v ", got, want, numbers)
		}
	})


}

//reflect.DeepEqual is useful for seeing if any two variables are the same 
// not also typesafe
func TestSumAll(t *testing.T){
      got := SumAll([]int{1,3}, []int{0, 9})
	  want := []int {3, 9}
    
	  if !reflect.DeepEqual(got, want) {
		   t.Errorf("got %v want %v", got, want)
	  }


}


//SUmalltails 

func TestSumAllTails(t *testing.T){
	

	t.Run("Make sum of slices ", func (t *testing.T){
        got := SumAllTails([]int{1,2}, []int{0,9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
	})

	t.Run("Safely sum empty slice", func(t *testing.T){
       got := SumAllTails([]int{}, []int{3,4,5})
	   want := []int{0, 9}
       
      if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)

	}})
}















