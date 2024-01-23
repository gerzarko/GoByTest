package concurrency

import (
	"fmt"
    "time"
)



func ConcurrencyMe (){

    start := time.Now()
    defer func ()  {
      fmt.Println(time.Since(start))  
    }()

    channel := make(chan string,2)
    go func(){



    sayHiCon(channel,2,"german")        
    sayHiCon(channel,2,"ivo")        
    }()
    for i := 0;i < 2; i++{
        fmt.Println(<- channel) 
    }

}


func sayHiCon(channel chan string, rounds int,name string)  {
        channel <-sayHi(name) 
        
}




func WIthoutConc (){

    start := time.Now()
    defer func ()  {
      fmt.Println(time.Since(start))  
        }()

    var all [2]string

    all[0] = sayHi("german")
    all[1] = sayHi("ivo")


}


func sayHi (name string)string{

    return fmt.Sprint(name)

}
