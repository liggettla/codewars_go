package main

import("fmt";"math")

//function that classifies an integer as either even or odd
func EvenOrOdd(number int) (status string) {
    if int(math.Abs(float64(number))) % 2 == 0 {
        status = "Even"
    } else if int(math.Abs(float64(number))) % 2 == 1 {
        status = "Odd"
    }
    return
}

func main() {
    x := 6
    status := EvenOrOdd(x)
    fmt.Println(status)

    y := -5
    status = EvenOrOdd(y)
    fmt.Println(status)
}
