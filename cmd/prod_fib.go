package main

import("fmt")

// returns the fibonacci sequence in the form of an array up to a maximum value
func fibSeq(max uint64) []uint64 {
    seq := []uint64{0,1,1}
    var end uint64 = 1

    for end <= max {
        end = (seq[len(seq)-1] + seq[len(seq)-2])
        if end < max { seq = append(seq, end) }
    }
    return seq
}

// checks if any product = n * (n+1) pairs exist within
// the fibonacci sequence
func ProductFib(product uint64) [3]uint64 {
    seq := fibSeq(product)
    for i := 1; i < len(seq); i++ {
        if product == seq[i] * seq[i+1] {
            return [3]uint64{seq[i], seq[i+1], 1}
        }
    }
    return [3]uint64{seq[len(seq)-2], seq[len(seq)-1], 0}
}

func main() {
    var product uint64 = 74049690 
    seq := fibSeq(product)
    fmt.Println(seq)
    answer := ProductFib(product)
    fmt.Println(answer)
}
