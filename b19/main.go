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

	const VALUE_MAX = 1000 * 100
	twoDim := make([][]int, n+1)
	twoDim[0] = make([]int, VALUE_MAX)
	for i := 1; i < VALUE_MAX; i++ {
		twoDim[0][i] = -1
	}

	for i := 1; i < n+1; i++ {
		twoDim[i] = make([]int, VALUE_MAX)
		for j := 0; j < VALUE_MAX; j++ {
			twoDim[i][j] = -1
			if j < v[i] || twoDim[i-1][j-v[i]] < 0 || twoDim[i-1][j-v[i]]+w[i] > knapsack {
				// 選ばない場合
				twoDim[i][j] = twoDim[i-1][j]
				continue
			}
			twoDim[i][j] = max(twoDim[i-1][j], twoDim[i-1][j-v[i]]+w[i])
		}
	}
	maxVal := 0
	for i, v := range twoDim[n] {
		if v > 0 {
			maxVal = i
		}
	}
	fmt.Println(maxVal)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
