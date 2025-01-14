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

	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
