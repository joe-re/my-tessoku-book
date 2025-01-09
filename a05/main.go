package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nstr, kstr string
	fmt.Scan(&nstr)
	fmt.Scan(&kstr)
	n, _ := strconv.Atoi(nstr)
	k, _ := strconv.Atoi(kstr)
	count := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			diff := k - (i + j)
			if diff > 0 && diff <= n {
				count++
			}
		}
	}
	fmt.Println(count)
}
