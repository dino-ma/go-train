package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

var a int = 1
var b int = 2

func init()  {
	Pi = 4 * math.Atan(1)
}

var Pi float64

func main() {


	fmt.Println(Pi)

	var goods string = runtime.GOOS
	fmt.Printf("This operting system is: %s\n", goods)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	var c = 3
	var d = 4
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	fmt.Println(&d)

	e := 5
	f := 6
	fmt.Println(&e)
	fmt.Println(&f)


	var vara string = "abc"
	fmt.Println(vara)

	ax, bx, cx := 5, 6, "asdf"

	fmt.Println(ax)
	fmt.Println(bx)
	fmt.Println(cx)
}
