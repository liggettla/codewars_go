/*
Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of n^3, the cube above will have volume of (n-1)^3 and so on until the top which will have a volume of 1^3.

You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?

The parameter of the function findNb (find_nb, find-nb, findNb) will be an integer m and you have to return the integer n such as n^3 + (n-1)^3 + ... + 1^3 = m if such a n exists or -1 if there is no such n.
*/

package main

import("fmt";"math")

func findNb(m int) (int) {
    var volume float64 = 0
    var cubes int
    for i:=0; volume<float64(m); i++ {
        volume = volume + math.Pow(float64(i),3)
        cubes = i
    }
    if volume == float64(m) {
        return cubes
    } else {
        return -1
    }
}

func main() {
    num_cubes := findNb(1071225)
    fmt.Println(num_cubes)
    num_cubes = findNb(91716553919377)
    fmt.Println(num_cubes)
}
