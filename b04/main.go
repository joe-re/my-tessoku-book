package main

import (
	"fmt"
)

func main() {
	var binary_str string
	fmt.Scan(&binary_str)
	var result = 0
	for i := len(binary_str) - 1; i >= 0; i-- {
		bit_value := int(binary_str[i] - '0')
		result += bit_value * (1 << (len(binary_str) - (i + 1)))
	}
	fmt.Println(result)
}
