package main

import (
	"bytes"
	"fmt"
	"main/Mocking"
)


func main() {


    buffer := &bytes.Buffer{}
    Mocking.Countdown(buffer)
    fmt.Println(buffer.String())
}
