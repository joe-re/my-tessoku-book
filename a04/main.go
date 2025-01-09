package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nstr string
	fmt.Scan(&nstr)
	var n, _ = strconv.Atoi(nstr)
	var result = ""
	for i := 9; i >= 0; i-- {
		result += strconv.Itoa((n / (1 << i)) % 2)
	}
	fmt.Println(result)
}
