package main

import "testing"

/* func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}


} */

func TestHello(t *testing.T) {
	got := Hello("Chris", "")
	want := "Hello, Chris"

	if got != want {
       t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T){

    assertCorrectMessage := func (t testing.TB, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want )
		}
	}

	t.Run("Saying hello to people", func(t *testing.T){
		got:= Hello("Chris", "")
		want:= "Hello, Chris"
		
		assertCorrectMessage(t, got, want)
	})


	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T){
            got := Hello("","")
			want := "Hello, World"

			assertCorrectMessage(t, got, want)
	})

	t.Run("Add language to the greeting", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Holla, Elodie"

			assertCorrectMessage(t, got, want)

	})


	t.Run("Say hello in  French", func(t *testing.T){
		got := Hello("Fred","French")
		want := "Bonjour, Fred"
		assertCorrectMessage(t, got, want)
      
	})
}
