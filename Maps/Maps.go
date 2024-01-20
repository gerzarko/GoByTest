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
    return "",errors.New("there is no key with that name")

}

func (d Dictionary) Add (key string,value string){
    d[key] = value

}



