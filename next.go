package main

import (
	"fmt"
)

func grade(score int) string  {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic("lalalal")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	//const filename  = "ab阿斯顿发c.txt"
	//if contents, err := ioutil.ReadFile(filename); err == nil{
	//	fmt.Println(string(contents))
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("canot print file content:", err)
	//	fmt.Printf("%s\n", contents)
	//}
	fmt.Println(
		grade(0),
		grade(59),
		grade(12),
		grade(66),
		grade(77),
		grade(88),
		grade(9),
		grade(1))


}
