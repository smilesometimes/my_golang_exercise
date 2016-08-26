package main

import (
	"fmt"
	"my_golang_exercise/go_stl/algorithm"
)

func main() {
	s1 := []int{1, 2, 4, 5, 9}
	s2 := []int{2, 6, 7, 8}
	l := len(s1) + len(s2)
	d := make([]int, l, l)
	algorithm.Merge(s1, s2, d)

	fmt.Println(d)
}
