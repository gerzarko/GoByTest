package Maps

import "errors"

// import "fmt"

type Dictionary map[string]string


func (d Dictionary) Search(toSearch string)(string,error){
    for k,v := range d{
       if k == toSearch {
           return v,nil
       }

    }
    return "",errors.New("could not find the word you were looking for")

}

func (d Dictionary) Add (key string,value string)error{
    d[key] = value
    return nil


}



