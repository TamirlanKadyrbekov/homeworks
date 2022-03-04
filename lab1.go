// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions

	fmt.Println(Fizzbuzz(6))
	fmt.Println(Fizzbuzz(10))
	fmt.Println(Fizzbuzz(15))
	fmt.Println(IsPrime(121))
	fmt.Println(IsPrime(71))
	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("car"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true

}
