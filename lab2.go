// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("ParsePhone() test")
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	fmt.Println("Anagram() test")
	fmt.Println(Anagram("12345", "52314"))
	fmt.Println(Anagram("21435", "53241"))
	fmt.Println(Anagram("12346", "52314"))
	fmt.Println(Anagram("123456", "52314"))
	fmt.Println("FindEvens() test")
	fmt.Println(FindEvens([]int{1, 2, 3, 4}))
	fmt.Println("SliceProduct() test")
	fmt.Println(SliceProduct([]int{1, 2, 3}))
	fmt.Println("Unique() test")
	fmt.Println(Unique([]int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8}))
	fmt.Println("InvertMap() test")
	fmt.Println(InvertMap(map[string]int{"A": 1, "B": 2, "C": 3}))
	fmt.Println("TopCharacters() test")
	fmt.Println(TopCharacters("123123122", 2))
	fmt.Println(TopCharacters("hello world", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {

	digits := make([]byte, 10)
	index := 0
	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			digits[index] = byte(ch)
			index++
		}
	}

	fst := string(digits[:3])
	mid := string(digits[3:6])
	end := string(digits[6:10])
	return fmt.Sprintf("(%v) %v-%v", fst, mid, end)
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	for _, ch := range s1 {
		m1[ch]++
	}
	m2 := make(map[rune]int)
	for _, ch := range s2 {
		m2[ch]++
	}
	for ch, cnt := range m1 {
		if cnt != m2[ch] {
			return false
		}
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(nums []int) []int {

	var evens []int
	for _, i := range nums {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(nums []int) int {

	product := 1
	for _, i := range nums {
		product *= i
	}
	return product
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {

	uniqueMap := make(map[int]bool)
	for _, i := range e {
		uniqueMap[i] = true
	}

	var uniques []int
	for i, _ := range uniqueMap {
		uniques = append(uniques, i)
	}

	return uniques
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	m := make(map[int]string)
	for k, v := range kv {
		m[v] = k
	}
	return m
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for r, cnt := range counts {
		if cnt <= k {
			delete(counts, r)
		}
	}
	return counts
}
