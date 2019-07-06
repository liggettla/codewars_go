package main

import("fmt")

// returns the fibonacci sequence in the form of an array up to a maximum value
func fibSeq(max uint64) []uint64 {
    seq := []uint64{0,1}
    var end uint64 = 1

    for end < max {
        end = (seq[len(seq)-1] + seq[len(seq)-2])
        if end < max { seq = append(seq, end) }
    }
    return seq
}

func main() {
    var max uint64 = 20
    seq := fibSeq(max)
    fmt.Println(seq)
}
