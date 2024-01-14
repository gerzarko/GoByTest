package Stringss

import (
	"strings"

)


func findTimesString (baseString string,searchString string)int{

    totalTimes := strings.Count(baseString,searchString)
    return totalTimes

}


func jointStrings(firstString string, secondString string) string {

    arrayOfStrings := []string {firstString,secondString} 
    stringTOReturn := strings.Join(arrayOfStrings," ")
    return stringTOReturn


}

func trimString (baseString string, trimString string) string {

    trimmedString := strings.Trim(baseString,trimString) 
    return trimmedString

}

func conversionToRune(baseString string)[]rune{
    runa := []rune(baseString)
    return runa 

}









