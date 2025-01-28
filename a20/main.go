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
	s := strings.Split(scanner.Text(), "")
	scanner.Scan()
	t := strings.Split(scanner.Text(), "")

	twoDim := make([][]int, len(s)+1)
	twoDim[0] = make([]int, len(t)+1)

	for i := 1; i < len(s)+1; i++ {
		twoDim[i] = make([]int, len(t)+1)
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				twoDim[i][j] = twoDim[i-1][j-1] + 1
			} else {
				twoDim[i][j] = max(twoDim[i][j-1], twoDim[i-1][j])
			}
		}
	}
	fmt.Println(twoDim[len(s)][len(t)])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
