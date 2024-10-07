package main

import (
	"fmt"
)

func main() {
	var n int
    fmt.Print("Enter number of elements in slice: ")
    fmt.Scan(&n)

    numbersSlice := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter element %d: ", i+1)
        fmt.Scan(&numbersSlice[i])
    }

	output:=findSecondLargest(numbersSlice)
	if output==-1{
		fmt.Println("There is no second largest element") 
	} else{
		
		fmt.Println("Second Largest is:",output)

	}


}

func findSecondLargest(numSlice []int) int {

	//Assuming input slice only has whole numbers in it

	largest := -1
	secondLargest := -1

	for _, value := range numSlice {

		if largest < value {
			secondLargest = largest
			largest = value

		} else if value > secondLargest && value < largest {
			secondLargest = value
		}
	}


	return secondLargest
}