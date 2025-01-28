package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := strings.Split(" "+scanner.Text(), "")

	scanner.Scan()
	t := strings.Split(" "+scanner.Text(), "")

	twoDim := make([][]int, len(s))
	twoDim[0] = make([]int, len(t))

	for i := 0; i < len(s); i++ {
		twoDim[i] = make([]int, len(t))
		for j := 0; j < len(t); j++ {
			if i == 0 {
				// 1行目は必ず=j
				twoDim[i][j] = j
				continue
			}
			if j == 0 {
				// 1列目は必ず=i
				twoDim[i][j] = i
				continue
			}

			c := 1
			if s[i] == t[j] {
				c--
			}
			above := twoDim[i-1][j] + 1
			left := twoDim[i][j-1] + 1
			aboveLeft := twoDim[i-1][j-1] + c
			twoDim[i][j] = min(above, left, aboveLeft)
		}
	}
	fmt.Println(twoDim[len(s)-1][len(t)-1])
}

func min(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}
