package main

import (
	"fmt"
)

func printSlice(s []int)  {
	fmt.Println("len=%d, cap=%x\n", s, len(s), cap(s))
}

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}

	s2 := make([]int, 16)

	//s3 := make([]int, 10, 32)
	//printSlice(s1)
	//printSlice(s2)
	//printSlice(s3)

	copy(s2, s1)
	printSlice(s2)
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	fmt.Println("poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2)
	//copy(s1, s2)
	//printSlice(s1)


}
