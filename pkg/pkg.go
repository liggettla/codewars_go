package pkg

import("fmt")

func Test() {
    fmt.Println("Import Successful")
}

// reverses the char order of a string
func StrRev(x string) string {
    var rev string
    for i:=len(x)-1; i>=0; i-- {
        rev += string(x[i])
    }
    return rev
}

// returns a list of all prime numbers less than max
func Primes(max int) []int {
    // list of all primes that we find
    var primes []int
    primes = append(primes, 2)

    var is_prime = true
    for i:=3; i<=max; i++ {
        is_prime = false
        // only odds can be prime
        if i%2 != 0 {
            is_prime = true
            // should not be divisible by any previous
            // prime numbers
            for _, x := range primes {
                if i%x == 0 {
                    is_prime = false
                    break
                }
            }
        }
        if is_prime {
            primes = append(primes, i)
        }
    }
    return primes
}

// finds all permutations of list items
func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

// non-recursive heap's algorithm
/*
procedure generate(n : integer, A : array of any):
    //c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k+1, A) is called
    c : array of int

    for i := 0; i < n; i += 1 do
        c[i] := 0
    end for

    output(A)
    
    //i acts similarly to the stack pointer
    i := 0;
    while i < n do
        if  c[i] < i then
            if i is even then
                swap(A[0], A[i])
            else
                swap(A[c[i]], A[i])
            end if
            output(A)
            //Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
            c[i] += 1
            //Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
            i := 0
        else
            //Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
            c[i] := 0
            i += 1
        end if
    end while
*/

