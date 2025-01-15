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

	if !twoDim[n][s] {
		fmt.Println(-1)
		os.Exit(0)
	}

	cards := make([]int, 0)
	curent := s
	for i := n; i > 0; i-- {
		for j := curent; j > 0; j-- {
			if twoDim[i-1][j] {
				// カードを選ばなかった
				break
			}
			if j-a[i] >= 0 && twoDim[i-1][j-a[i]] {
				// カードを選んでいる
				cards = append(cards, i)
				curent -= a[i]
				break
			}
		}
	}
	fmt.Println(len(cards))
	for i := len(cards) - 1; i >= 0; i-- {
		fmt.Print(cards[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
}
