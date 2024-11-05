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

}




func assertError(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
