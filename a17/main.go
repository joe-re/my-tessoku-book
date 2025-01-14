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
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, n-2)
	for i := 0; i < n-2; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = a[0]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+a[i-1], dp[i-2]+b[i-2])
	}

	routes := make([]int, 0)
	for i := n - 1; i > 0; {
		if dp[i] == dp[i-1]+a[i-1] {
			routes = append(routes, i)
			i--
		} else if dp[i] == dp[i-2]+b[i-2] {
			routes = append(routes, i-1)
			i -= 2
		}
	}
	fmt.Println(len(routes) + 1)
	for i := len(routes) - 1; i >= 0; i-- {
		fmt.Print(strconv.Itoa(routes[i]) + " ")
	}
	fmt.Print(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
