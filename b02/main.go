package main

import (
	"fmt"
	"strconv"
)

func main() {
	var astr string
	var bstr string
	fmt.Scan(&astr)
	fmt.Scan(&bstr)

	var a, _ = strconv.Atoi(astr)
	var b, _ = strconv.Atoi(bstr)

	var found = false
	for i := a; i <= b; i++ {
		if 100%i == 0 {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
