package main

import "fmt"

func square_sum() int {
    total := 0
    for i:=1; i<=100; i++ {
        total += i
    }

    return total * total
}

func sum_square(x int) int {
    if x == 1 {
        return 1
    }
    return x * x + sum_square(x-1)
}

func main() {
    x := square_sum()

    num := 100
    y := sum_square(num)
    fmt.Println(x-y)

}
