package Mocking

import (
	"bytes"
	"fmt"
	"time"
)

func Countdown(out*bytes.Buffer){

    for i:=3; i > 0;i--{

        fmt.Println(out,i) 
        time.Sleep(1 * time.Second)
    }

    fmt.Fprintf(out,"Go!") 
}

