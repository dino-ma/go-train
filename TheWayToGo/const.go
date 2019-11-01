package main

import (
	"fmt"
	"math"
)

func main() {

	var n int16 = 34
	var m int32
	m = int32(n)

	fmt.Println(m)
	fmt.Println(n)

	//num, error = Uint8FromInt(n)
	//fmt.Println(num)
	//fmt.Println(error)
}

func Uint8FromInt(n int) (unit8, error) {
	if 0<=n && n <= math.MaxUint8{
		return unit8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the unit8 range", n)
}