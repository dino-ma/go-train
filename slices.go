package main

import "fmt"

func updateSlices(s []int)  {
	s[0] = 100

}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])


	s1 := arr[2:]
	fmt.Println("=========")
	updateSlices(s1)
	fmt.Println("s1", s1)
	s2 := arr[:]
	updateSlices(s2)
	fmt.Println("s2", s2)


	a :=[...]int{1,2,3,4,5,6,}
	fmt.Println(a)

	b :=[...]int{1,2,3,4}

	a = append(a, b)
	fmt.Println(a)


}
