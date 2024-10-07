package main

import (
    "fmt"
    "time"
)

func main() {
    hour := time.Now().Hour()
    minute := time.Now().Minute()

    switch {
    case hour >= 6 && hour < 11:
        fmt.Println("Good morning")
    case hour == 11 && minute == 0:
        fmt.Println("Good morning")
    case hour >= 11 && hour < 16:
        fmt.Println("Good afternoon")
    case hour == 16 && minute == 0:
        fmt.Println("Good afternoon")
    case hour >= 16 && hour < 21:
        fmt.Println("Good evening")
    case hour == 21 && minute == 0:
        fmt.Println("Good evening")
    default:
        fmt.Println("Good night!")
    }
}
