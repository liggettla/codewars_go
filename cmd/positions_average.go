/*
Suppose you have 4 numbers: '0', '9', '6', '4' and 3 strings composed with them:

s1 = "6900690040"
s2 = "4690606946"
s3 = "9990494604"
Compare s1 and s2 to see how many positions they have in common: 0 at index 3, 6 at index 4, 4 at index 8 ie 3 common positions out of ten.

Compare s1 and s3 to see how many positions they have in common: 9 at index 1, 0 at index 3, 9 at index 5 ie 3 common positions out of ten.

Compare s2 and s3. We find 2 common positions out of ten.

So for the 3 strings we have 8 common positions out of 30 ie 0.2666... or 26.666...%

Given n substrings (n >= 2) in a string s our function pos_average will calculate the average percentage of positions that are the same between the (n * (n-1)) / 2 sets of substrings taken amongst the given n substrings. It can happen that some substrings are duplicate but since their ranks are not the same in s they are considered as different substrings.

The function returns the percentage formatted as a float with 10 decimals but the result is tested at 1e.-9 (see function assertFuzzy in the tests).

Example:
Given string s = "444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490" composing a set of n = 10 substrings (hence 45 combinations), pos_average returns 29.2592592593.

In a set the n substrings will have the same length ( > 0 ).
*/

package main

import ("fmt";"strings")

func split_up(str1 string, str2 string) ([]string, []string) {
    return strings.Split(str1, ""), strings.Split(str2, "")
}

func compare_strings(str1 []string, str2[]string) int {
    var count int = 0
    for i := range str1 {
        if str1[i] == str2[i] { count++ }
    }
    return count
}

func array_combinations(a []string) {
    combinations := []Tuple
}

type Tuple struct {
    a, b int
}

// returns the percent conserved chars
func PosAverage(s string) []string {
    return strings.Split(s, ", ")
}

func main() {
    s := "444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490"
    percent_overlap := PosAverage(s)
    fmt.Println(percent_overlap)
    for i:=range percent_overlap { fmt.Println(percent_overlap[i]) }
}
