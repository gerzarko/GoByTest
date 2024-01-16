package ArraySlices
import "fmt"

func Sum(arrayOfNumber []int) int {

	total := 0
	for index, number := range arrayOfNumber {
		fmt.Print(index)
		total += number
	}
	return total
}

func SumAll(firstSlice []int, secondSlice []int) []int {
    firstTotal := Sum(firstSlice)
    secondTotal := Sum(secondSlice)
	returnSlice := []int{firstTotal,secondTotal}


	return returnSlice

}
func SumAllTails(firstArray []int,secondArray[]int)[]int{
    ultimoDelPrimero := firstArray[len(firstArray)-1]
    ultimoSegundo := secondArray[len(secondArray) -1 ]
    retorno := []int{ultimoDelPrimero,ultimoSegundo}
    return retorno  
 
}
