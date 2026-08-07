package main

import "fmt"

// isPrime returns true if n is a prime number
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("%d prime? %v\n", n, isPrime(n))
}
