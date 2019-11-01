package main

import "fmt"

const  (
	MSJ_A = iota
	MSJ_B = 222
	MSJ_C
	MSJ_D
)

func main() {

	fmt.Println(MSJ_A)
	fmt.Println(MSJ_B)
	fmt.Println(MSJ_C)
	fmt.Println(MSJ_D)


	var a int
	var b int
	a = 15
	b = b + 5
	b = a + a
	fmt.Println(a)
	fmt.Println(b)



	var n int16 = 34
	var m int32
	m = int32(n)
	fmt.Println(n)
	fmt.Println(m)
}
