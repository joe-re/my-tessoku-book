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

	scanner.Scan()
	knapsack, _ := strconv.Atoi(scanner.Text())

	w := make([]int, n+1)
	v := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		scanner.Scan()
		w[i], _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}

	twoDim := make([][]int, n+1)
	twoDim[0] = make([]int, knapsack+1)
	for i := range twoDim[0] {
		twoDim[0][i] = -1
	}
	twoDim[0][0] = 0

	for i := 1; i < n+1; i++ {
		twoDim[i] = make([]int, knapsack+1)
		for j := 0; j <= knapsack; j++ {
			twoDim[i][j] = -1
			if j < w[i] || twoDim[i-1][j-w[i]] < 0 {
				// 何も選ばない
				twoDim[i][j] = twoDim[i-1][j]
				continue
			}
			twoDim[i][j] = max(twoDim[i-1][j], twoDim[i-1][j-w[i]]+v[i])
		}
	}
	value := 0
	for _, v := range twoDim[n] {
		value = max(value, v)
	}
	fmt.Println(value)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
