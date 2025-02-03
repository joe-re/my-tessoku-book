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

	a := make([]int, n-1)
	b := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	// fmt.Println(n)
	// fmt.Println(a)
	// fmt.Println(b)

	routes := make([]int, n)
	for i := 0; i < n; i++ {
		routes[i] = -1000
	}
	routes[0] = 0

	for i := 0; i < n-1; i++ {
		ax := a[i]
		bx := b[i]
		routes[ax-1] = max(100+routes[i], routes[ax-1])
		routes[bx-1] = max(150+routes[i], routes[bx-1])
	}
	// fmt.Println(routes)
	fmt.Println(routes[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
