package main

import (
	"fmt"
)

const c  = "C"

var v int = 5

type T struct {}

type Msj struct {
}

func init() {
	fmt.Printf("this func is init\n")
}

const msj float64 = 1.234

const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	iotaaa = 1
	iotaa = iota
	iotab = iota
	_MSJ_CONST = 2
)

func main() {
	//var a int
	//Func1()
	//fmt.Println(a)
	//fmt.Println(v)
	//fmt.Println(c)
	//z := T{}
	//z.Method1()
	//fmt.Println(msj)
	//
	//n := Msj {}
	//n.Method1()
	//
	//fmt.Println(Unknown)
	//fmt.Println(Female)
	//fmt.Println(Male)

	//fmt.Println(iotaaa )
	//
	//fmt.Println(iotaa)
	//fmt.Println(iotab)
	//fmt.Println(_MSJ_CONST)

	if 1==1 {
		fmt.Println("ok")
	}
	fmt.Println(1 << 10 << 10 >> 10)

	fmt.Println(6 &^ 3)

	a := 3
	if a > 0 && (10/a) > 1 {
		fmt.Println("10/a OK")
	}
}

func (t T) Method1()  {
	a := 655321
	b := int(a)
	fmt.Println(b)

	//..
}

func (x Msj) Method1() {
	fmt.Println("Msj Struct")
}

func Func1()  {
	fmt.Println("this func is Func1\n")
}