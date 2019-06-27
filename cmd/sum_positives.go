package main

import("fmt")

// sum all the positive integers in an array
func sum_positives(nums []int) int {
    sum := 0
    for _, num := range nums {
        if num > 0 {
            sum += num
        }
    }
    return sum
}

func main() {
    test := []int{1,-4,7,12}
    sum := sum_positives(test)
    fmt.Println(sum)
}
