package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nstr string
	fmt.Scan(&nstr)
	var n, _ = strconv.Atoi(nstr)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		var vstr string
		fmt.Scan(&vstr)
		var v, _ = strconv.Atoi(vstr)
		a[i] = v
	}
	var found = false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if i == j {
				continue
			}
			for k := 0; k < n && !found; k++ {
				if k == i || k == j {
					continue
				}
				if a[i]+a[j]+a[k] == 1000 {
					found = true
					break
				}
			}
		}
	}
	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
