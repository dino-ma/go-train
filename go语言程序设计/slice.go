package main

import (
	"fmt"
	"runtime"
)

const B float64  = 1 << (iota * 10)


func main() {


	fmt.Println(B)

	fmt.Printf(runtime.Version())
	//if len(os.Args) == 1 {
	//	fmt.Printf("usage: %s <whole-number>\n ", filepath.Base(os.Args[0]))
	//	os.Exit(2)
	//}
	//stringOfdigits := os.Args[1]
	//for row := range bigDigits[0] {
	//	line := " "
	//	for column := range stringOfdigits {
	//		digit := stringOfdigits[column] - '0'
	//		if 0 <= digit && digit <=9 {
	//			line += bigDigits[digit][row] + " "
	//
	//		} else {
	//			log.Fatal(" invalid whole number")
	//		}
	//
	//	}
	//	fmt.Println(line)
	//}

}
