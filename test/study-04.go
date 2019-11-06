package main

import "fmt"

const a int = 1
const b string = "A"
const c, d, e = 1, 2, 3
const (
	f = 1222
	g
	h
)

const (
	A = iota
	C
	D
	E = iota
	F
)

const (
	_         = iota
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(A)
	fmt.Println(F)

	fmt.Println(1 << 10 << 10 >> 10)
	fmt.Println(1 == 1)
	fmt.Println(6 &^ 14)

	a := 1
	if (a > 0) && (10/a) > 1 {
		fmt.Println("ok")
	}

	//homework
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
