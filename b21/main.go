package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := strings.Split(scanner.Text(), "")

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for len := 2; len <= n; len++ {
		for l := 0; l < n-len; l++ {
			r := len + l
			// fmt.Print("-", l, r, "-")
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1] + 2
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}
		}
		// fmt.Println()
	}

	fmt.Println(dp[0][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
