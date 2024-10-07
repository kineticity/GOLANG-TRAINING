package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func fibonnacciSum(number int) int {
	if number <= 1 {
		return 0
	}
	a, b := 0, 1
	sum := a + b
	for i := 2; i < number; i++ {
		next := a + b
		sum += next
		a = b
		b = next
	}

	return sum
}

func countOddEvenZero(arr []int) (int, int, int) {
	zeroCount, evenCount, oddCount := 0, 0, 0
	for _, num := range arr {
		if num == 0 {
			zeroCount++
		} else if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return oddCount, evenCount, zeroCount
}

func main() {
	var number int
	fmt.Println("Enter number to check if prime/get fibonnacci sum: ")
	fmt.Scan(&number)

	p := isPrime(number)
	fmt.Printf("Is %d Prime?:%v\n", number, p)

	sum := fibonnacciSum(number)
	fmt.Println("Sum:", sum)

	var n int
	fmt.Println("Enter number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	oddCount, evenCount, zeroCount := countOddEvenZero(arr)

	fmt.Printf("Odd elements count: %d\n", oddCount)
	fmt.Printf("Even elements count: %d\n", evenCount)
	fmt.Printf("Zero elements count: %d\n", zeroCount)

}
