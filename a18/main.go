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
	s, _ := strconv.Atoi(scanner.Text())

	// 最初は0にする
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	// 0枚目位の0は必ずtrue
	twoDim := make([][]bool, n+1)
	twoDim[0] = make([]bool, s+1)
	twoDim[0][0] = true

	for i := 1; i <= n; i++ {
		twoDim[i] = make([]bool, s+1)
		for j := 0; j <= s; j++ {
			if twoDim[i-1][j] == true {
				// カードを選ばないことができるのでtrue
				twoDim[i][j] = true
				continue
			}
			if a[i] <= j && twoDim[i-1][j-a[i]] {
				// カードを選んだ結果jになるのでtrue
				twoDim[i][j] = true
			}
		}
	}

	if twoDim[n][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
