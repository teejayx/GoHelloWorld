package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//& pointer let us point to the address in memory of a value 
//* pointer let us point to the value instead of copy 

type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type DefaultSleeper struct{}
func (d *DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
}

const (finalWorld = "Go!"
    countdownStart = 3
 )

func main (){
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}



 func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i--{
	time.Sleep(1 * time.Second)
	fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWorld)
 }
	 