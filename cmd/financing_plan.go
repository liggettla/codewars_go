package main

import ("fmt")

func Finance(n int) int {
    days := n+1
    max := n
    var start int = 0
    var total int = 0
    for {
        if days == 0 { break }
        for i:=start; i<=max; i++ {
            total += i
        }
        days -= 1
        start += 2
        max += 1
    }
    return total
}

// not my solution, I couldn't find
// this pattern
func Better_Finance(n int) int {
    return n * (n+1) * (n+2) / 2
    }

func main() {
    total := Finance(5000)
    fmt.Println(total)
}
