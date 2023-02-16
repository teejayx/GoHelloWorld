package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

//racer without select
func Racer(a, b string) (winner string){
	aDuration := measureResponseTime(a)

	bDuration := measureResponseTime(b)

	if aDuration < bDuration{
		return a
	}

    return b
}


func measureResponseTime(url string ) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}



// racer with select 

func Racer2 (a, b string, timeout time.Duration) (winner string, error error ){
	select {
	case <-ping(a):
		      return a, nil
	case <-ping(b): 
		      return b, nil
	case <- time.After(timeout): {
		return "", fmt.Errorf("timed out after wwaiting for %s and %s", a, b)
	}
	}
}


func ping(url string) chan struct{}{
	//always use make channel to creare a new chan so it's not nil
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}
