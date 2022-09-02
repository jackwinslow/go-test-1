package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	fmt.Print("Please enter an integer: ")
	// Scan input from command line
	fmt.Scan(&x)

	// Array
	// arr_start := time.Now()
	// var a[x + 1]int
	// for i := 0; i <= x; i++ {
	// 	a[i] = i
	// }
	// arr_end := time.Since(arr_start)

	// Slice
	// Start timer
	slice_start := time.Now()
	// Init slice
	s := make([]int, 0)
	// Add values
	for i := 0; i <= x; i++ {
		s = append(s, i)
	}
	// End timer
	slice_end := time.Since((slice_start))

	// Map
	// Start timer
	map_start := time.Now()
	// Init map
	var m = map[int]int{}
	// Add values
	for i := 0; i <= x; i++ {
		m[i] = i
	}
	// End timer
	map_end := time.Since((map_start))

	// First output
	fmt.Println(slice_end, map_end)

	// Array 2

	// Slice 2
	// Start timer
	slice_start2 := time.Now()
	// Increment values
	for i := 0; i <= x; i++ {
		s[i] = i + 1
	}
	// End timer
	slice_end2 := time.Since((slice_start2))

	// Map 2
	// Start timer
	map_start2 := time.Now()
	// Increment values
	for i := 0; i <= x; i++ {
		m[i] = i + 1
	}
	// End timer
	map_end2 := time.Since((map_start2))

	// Second output
	fmt.Println(slice_end2, map_end2)
}
