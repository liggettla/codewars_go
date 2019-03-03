/*
Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.

Each two non empty parts will be in a pair (or an array for languages without tuples or a structin C - C: see Examples test Cases - )
Each part will be in a string
Elements of a pair must be in the same order as in the original array.
Examples of returns in different languages:
a = ["az", "toto", "picaro", "zone", "kiwi"]
returns:
[["az", "toto picaro zone kiwi"], ["az toto", "picaro zone kiwi"], ["az toto picaro", "zone kiwi"], ["az toto picaro zone", "kiwi"]]
*/

package main

import("fmt";"strings")

func PartList(arr []string) string {
    var m,n string
    var final string = ""

    for i:=1; i<len(arr); i++ {
        m = strings.Join(arr[0:i], " ")
        n = strings.Join(arr[i:len(arr)], " ")
        final += "(" + m + ", " + n + ")"
    }
    return final
}

func main() {
    var parts = []string{"I","wish","I","hadn't","come"}
    answer := PartList(parts)
    fmt.Println(answer)
}
