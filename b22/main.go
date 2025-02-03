package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := 1; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, n)
	for i := 2; i < n; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	// fmt.Println(a)
	// fmt.Println(b)

	routes := make([]int, n)
	for i := 0; i < n; i++ {
		routes[i] = 1000000000
	}
	routes[0] = 0

	for i := 0; i < n; i++ {
		if i+1 < n {
			ax := a[i+1]
			routes[i+1] = min(routes[i]+ax, routes[i+1])
		}
		if i+2 < n {
			bx := b[i+2]
			routes[i+2] = min(routes[i]+bx, routes[i+2])
		}
	}
	// fmt.Println(routes)
	fmt.Println(routes[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
