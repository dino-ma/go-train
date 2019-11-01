package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loop() {
	sum := 0
	for i := 1; i<= 100; i++ {
		sum += i
		fmt.Println(sum, i)
	}
}

func coverToBin(n int) string  {
	reuslt := ""
	for ; n>0 ; n /= 2  {
		lsb := n % 2
		reuslt = strconv.Itoa(lsb) + reuslt
	}
	return reuslt
}

func printFile(filename string)  {
	file , err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for  {
		fmt.Println("abc")
	}
}

func main() {
	//loop()
	fmt.Println(
		coverToBin(5),
		coverToBin(13),
		coverToBin(15),
		coverToBin(1123123123123),
		coverToBin(0),
		)
	printFile("abc.txt")
	forever()
}
