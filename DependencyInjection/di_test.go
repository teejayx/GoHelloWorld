package main

import (
	"bytes"
	"testing"
)

//& goes in front of a variable when you want to get that variable's memory address.

//* goes in front of a variable that holds a memory address and resolves it

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got , want)
	}
}