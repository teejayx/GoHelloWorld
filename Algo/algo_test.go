package main

import (
	"testing"
    "reflect"
)


func TestSelectionSort(t *testing.T){
	got := InsertionSort([]int{5,2,4,6,1,3})
	want := []int {1,2,3,4,5,6}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
}