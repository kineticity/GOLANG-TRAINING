package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
    if number <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
        if number%i == 0 {
            return false
        }
    }
    return true
}

func findPrimeNumbersInRange(rangeStart int,rangeEnd int) []int {
    var sliceOfPrimes []int
    for i := rangeStart; i <= rangeEnd; i++ {
        if isPrime(i) {
            sliceOfPrimes = append(sliceOfPrimes, i)
        }
    }
    return sliceOfPrimes
}

func main() {
	var rangeStart int
    var rangeEnd int

	fmt.Print("Enter the start of the range: ")
    fmt.Scan(&rangeStart)
    fmt.Print("Enter the end of the range: ")
    fmt.Scan(&rangeEnd)

    primes := findPrimeNumbersInRange(rangeStart,rangeEnd)

    if len(primes) == 0 {
        fmt.Println("There are no prime numbers in given range")
    } else {
        fmt.Println("Prime numbers are:", primes)
    }
}
