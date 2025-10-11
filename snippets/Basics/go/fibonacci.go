// Metadata Header (MANDATORY)
// -----------------------------
// Program Title: Fibonacci Sequence (Iterative)
// Author: ghostrider2606
// Date: 2025-10-11
//
// Description: Generates the Fibonacci sequence up to the n-th term using an iterative approach.
//
// Language: Go
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// -----------------------------

package main

import "fmt"

// fibonacciIterative generates the first n Fibonacci numbers.
// Uses uint64 for large numbers.
func fibonacciIterative(n int) ([]uint64, error) {
    if n < 0 {
        return nil, fmt.Errorf("input must be a non-negative integer")
    }
    if n == 0 {
        return []uint64{}, nil
    }
    if n == 1 {
        return []uint64{0}, nil
    }

    sequence := make([]uint64, n)
    sequence[0] = 0
    sequence[1] = 1

    var a, b uint64 = 0, 1

    for i := 2; i < n; i++ {
        nextVal := a + b
        sequence[i] = nextVal
        
        // Shift values
        a = b
        b = nextVal
    }
    return sequence, nil
}

func main() {
    n_terms := 20
    result, err := fibonacciIterative(n_terms)
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Fibonacci sequence (first %d terms): %v\n", n_terms, result)
    }
}
