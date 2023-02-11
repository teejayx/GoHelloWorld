package main

import "testing"


func TestSearch(t *testing.T) {
 dictionary := Dictionary{"test": "this is just a test"}


 t.Run("Known word", func(t *testing.T) {
  got, _ := dictionary.Search("test")
  want := "this is just a test"

  assertString(t, got, want)

 })

 t.Run("Unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}
        assertError(t, err, ErrNotFound)

 })

 t.Run("Add Test", func( t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	dictionary.Add(word, definition)

     AssertDefnition(t, dictionary, word, definition)
 })

 t.Run("existing word", func(t *testing.T){
    word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word : definition}
    err := dictionary.Add(word, "new test")

	assertError(t, err, errWordExist)
	AssertDefnition(t , dictionary, word, definition)


 } )
 
}


func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want{
		t.Errorf("got error %q want %q", got, want)
	}
}

func AssertDefnition (t testing.TB, dicDictionary Dictionary, word, definition string){
	t.Helper()

	got, err := dicDictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}