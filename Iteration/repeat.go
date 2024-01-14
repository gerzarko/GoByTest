package Iteration

func Repeat(word string,numberToRepeat int8) string {

    for i := 0; i < int(numberToRepeat-1); i++ {
       word = word + "a"
    }	

    return word
}
