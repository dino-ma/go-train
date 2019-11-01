package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var msj int = 123
var dino string = "dino.ma"

func variableZeroValue() {
	var a,b,c int = 1,2,3
	var str string = "lalala"
	fmt.Println(a, b, c, str)
}

func variableInit()  {
	var a int
	var b string
	fmt.Println(a, b)
}

func variableTypeDeduction() {
	var a,v,c  = 3, 4, true
	fmt.Println(a, v, c)
}

//欧拉公式
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

//强制类型转换
func trigle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	const filename string = "abc.txt"
	const a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

func enums()  {
	const (
		cpp = iota
		php
		java
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, php, java, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInit()
	variableTypeDeduction()
	fmt.Println(msj, dino)
	euler()
	trigle()
	consts()
	enums()
}
