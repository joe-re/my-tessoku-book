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

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		c[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		d[i], _ = strconv.Atoi(scanner.Text())
	}

	MAX := 1501
	twoDimSum := make([][]int, MAX)
	for i := 0; i < MAX; i++ {
		twoDimSum[i] = make([]int, MAX)
	}
	for i := 0; i < n; i++ {
		twoDimSum[x[i]][y[i]] += 1
	}

	for i := 0; i < MAX; i++ {
		for j := 1; j < MAX; j++ {
			twoDimSum[i][j] += twoDimSum[i][j-1]
		}
	}

	for i := 1; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			twoDimSum[i][j] += twoDimSum[i-1][j]
		}
	}

	for i := 0; i < q; i++ {
		ans := twoDimSum[c[i]][d[i]] - twoDimSum[a[i]-1][d[i]] - twoDimSum[c[i]][b[i]-1] + twoDimSum[a[i]-1][b[i]-1]
		fmt.Println(ans)
	}
}
