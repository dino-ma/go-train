package main

import (
	"fmt"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b , nil
	case "-":
		return a - b , nil
	case "*":
		return a * b , nil
	case "/":
		return a / b , nil
	default:
		return 0, fmt.Errorf("error op")
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
	//q = a / b
	//r = a % b
	//return
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	//if result, err := eval(1, 5, "*"); err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Println(result)
	//}
	//fmt.Println(eval(1, 5, "123"))
	//q, _ :=div(13, 14)
	//fmt.Println(q )
	//var a int = 2
	//var pa *int = &a
	//var b *int = &a
	//fmt.Println(a)
	//*pa = 3
	//*b = 4

	a, b := 1111, 222
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
