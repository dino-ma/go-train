package main

import (
	"fmt"
)

type msjTestA struct {
	Name string
}

type msjTestB struct {
	Name string
}

type Tz int

func main() {
	a := msjTestA{}
	a.Print()
	b := msjTestB{}

	b.Print()


	//fmt.Println(a, b)

	//home work
	var c Tz
	c = 5
	c.Increase(100)
	fmt.Println(c)

}

func (tz * Tz) Increase(num int) {
	*tz += Tz(num)
}


func (msjTestA *msjTestA) Print() {
	msjTestA.Name = "马胜杰啦啦啦啦msjtestA"
	fmt.Println("msjtestA:", msjTestA)
}
func (msjTestB *msjTestB) Print() {
	msjTestB.Name = "msjtestb"
	fmt.Println("msjtestb:", msjTestB)
}
