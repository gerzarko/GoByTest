package main

import (
	// "bytes"
 //    "main/Mocking"
	// "fmt"
    "main/Concurrency"
)


func main() {


    // buffer := &bytes.Buffer{}
    // Mocking.Countdown(buffer)
    // fmt.Println(buffer.String())
    concurrency.ConcurrencyMe()
    concurrency.WIthoutConc()

}
