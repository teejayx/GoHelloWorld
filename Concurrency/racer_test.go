package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// use a package net/http/httptest to mock httpserver instead of making direct calls

//select  helps us synchronise processes really easy and readily 

func TestRacer(t *testing.T){
  // slowURL := "http://www.facebook.com"
   //fastUrl := "http://www.quii.co.uk"

t.Run("racer without conc", func(t *testing.T){

	slowServer := makeDelayedServer(20 * time.Millisecond)

	fastServer := makeDelayedServer(0 * time.Millisecond)
 
	//deferr calls the functoin at the end of the of the containing functionns 
   defer slowServer.Close()
   defer fastServer.Close()
 
 
  slowUrl := slowServer.URL
  fastUrl := fastServer.URL
 
	want := fastUrl
	got := Racer(slowUrl, fastUrl)
 
	if got != want {
	  t.Errorf("got %q, want %q", got, want)
	}
 

})

var tenSecondTimeout = 10 * time.Second
t.Run("racer without conc", func(t *testing.T){

	slowServer := makeDelayedServer(12 * time.Millisecond)

	fastServer := makeDelayedServer(13 * time.Millisecond)
 
	//deferr calls the functoin at the end of the of the containing functionns 
   defer slowServer.Close()
   defer fastServer.Close()

   //want := slowServer.URL
   //got := Racer(slowUrl, fastUrl)
 
     _, err := Racer2(slowServer.URL, fastServer.URL, tenSecondTimeout)

 
	if err != nil {
	  t.Errorf("expected an error but didnt get one")
	}
    

})



  

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
         time.Sleep(delay)
		w.WriteHeader((http.StatusOK))
	}))
   
}