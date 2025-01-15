package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	h := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(h[0] - h[1])
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+abs(h[i]-h[i-1]), dp[i-2]+abs(h[i]-h[i-2]))
	}

	routes := make([]int, 0)
	for i := n - 1; i >= 0; {
		routes = append(routes, i)
		if i-2 >= 0 && dp[i] == dp[i-2]+abs(h[i]-h[i-2]) {
			i -= 2
		} else {
			i--
		}
	}

	fmt.Println(len(routes))
	for i := len(routes) - 1; i >= 0; i-- {
		fmt.Print(routes[i] + 1)
		if i > 0 {
			fmt.Print(" ")
		}
	}

}
