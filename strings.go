package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	//fmt.Printf("%X\n ",[]byte(s))
	fmt.Println(s)
	for _, b := range []byte(s){
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	bytes := []byte(s)
	for len(bytes) >0  {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("%c ", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}


	//fmt.Println(len(s))

}
