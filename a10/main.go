package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	r := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		scanner.Scan()
		r[i], _ = strconv.Atoi(scanner.Text())
	}

	p := make([]int, n+1)
	p[1] = r[1]
	for i := 2; i <= n; i++ {
		p[i] = max(p[i-1], r[i])
	}
	q := make([]int, n+1)
	q[n] = r[n]
	for i := n - 1; i > 0; i-- {
		q[i] = max(q[i+1], r[i])
	}

	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	result := make([]int, d)
	for i := 0; i < d; i++ {
		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		e, _ := strconv.Atoi(scanner.Text())
		result[i] = max(p[s-1], q[e+1])
	}

	for i := 0; i < d; i++ {
		fmt.Println(result[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
