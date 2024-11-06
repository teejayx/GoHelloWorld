package main

import (
	"testing"
)


func TestStringManipulation (t *testing.T) {


     t.Run("Palindrome", func  (t *testing.T)  {
    
	text := "Madam"
	want := true
	got := Panlindrome(text)
    assertError(t, got, want)
} )


   t.Run("Anagram",  func(t *testing.T){
      word := "listen"
	  word2 := "silent"  
	  want := true
	  got := AnagramSort(word, word2)
       assertError(t, got, want)

} )
 t.Run("Anagram",  func(t *testing.T){
      word := "listen"
	  word2 := "silent"  
	  want := true
	  got := AnagramCountFreq(word, word2)
       assertError(t, got, want)

} )


 t.Run("`Reverse String",  func(t *testing.T){
      word := "listen"
	  want := "netsil"
	  got := ReversString(word)
       assertStringError(t, got, want)

} )

t.Run("Reserve String two Pointers", func(t *testing.T){
	  word := "listen"
	  want := "netsil"
	  got := ReversStringTwoPointer(word)
       assertStringError(t, got, want)
})

}




func assertError(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func assertStringError(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}