package main

import("fmt")
import("../pkg")

func main() {
    check := pkg.IsPalindrome("kayak")
    fmt.Println(check)
}
