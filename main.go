package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func partOne() {
	x := 0
	fmt.Print("Please enter an integer: ")
	// Scan input from command line
	fmt.Scan(&x)

	// Array
	// // Start timer
	// arr_start := time.Now()
	// // Init array
	// arr := [x]int
	// // Add values
	// for i := 0; i <= x; i++ {
	// 	arr[i] = i
	// }
	// // End timer
	// arr_end := time.Since(arr_start)
	// fmt.Println(arr)

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

func partTwo() {
	rand.Seed(time.Now().UnixNano())
	x := 0
	fmt.Print("Please enter an integer: ")
	// Scan input from command line
	fmt.Scan(&x)

	// Slice
	// Init slice
	s := make([]int, 0)
	// Add values
	for i := 0; i < x; i++ {
		s = append(s, rand.Intn(100))
	}
	// Making a copy of slice
	s2 := make([]int, len(s))
	copy(s2, s)

	// Start timer
	slice_start := time.Now()
	// Sort Slice
	s1 := sort.IntSlice(s)
	sort.Sort(s1)
	// End timer
	slice_end := time.Since((slice_start))

	// Start timer
	slice_start2 := time.Now()
	// Sort Slice
	s3 := sort.IntSlice(s2)
	sort.Stable(s3)
	// End timer
	slice_end2 := time.Since((slice_start2))

	// Slice Outputs
	fmt.Print(slice_end, slice_end2)
}

func main() {
	partOne()
	partTwo()
}
