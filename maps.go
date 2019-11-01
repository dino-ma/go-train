package main

import "fmt"

func lengthNeveryRepaetSubstr(str string) int {
	lastCurrent := make(map[byte]int)
	start := 0
	maxlength := 0
	for i, ch := range []byte(str)  {
		if lastI, ok := lastCurrent[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxlength {
			maxlength = i - start + 1
		}
		lastCurrent[ch] = i
	}

	return maxlength
}

func main() {
	m := map[string] string {
		"name" : "dino.ma",
		"course": "golang",
	}

	m2 := make(map[string]int)
	fmt.Println(m, m2)

}
