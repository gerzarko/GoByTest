package Mocking

import (
	"bytes"
	"fmt"
	"time"
)


type Sleeper interface{
    Sleep()
}

type SpySleeper struct{
    Calls int
}


func(s *SpySleeper) Sleep(){
    s.Calls++
}


func Countdown(out*bytes.Buffer,sleeper Sleeper){

    for i:=3; i > 0;i--{

        fmt.Fprintln(out,i) 
        time.Sleep(1 * time.Second)
        SpySleeper() 
    }

    fmt.Fprintf(out,"Go!") 
}

