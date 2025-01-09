package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(n int, printers []int, k int, sec int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += sec / printers[i]
	}
	return sum >= k
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	printers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		printers[i], _ = strconv.Atoi(scanner.Text())
	}

	MAX_SEC := 1_000_000_000_000

	l, r := 0, MAX_SEC

	result := 0

	for l <= r {
		m := (l + r) / 2
		if check(n, printers, k, m) {
			r = m - 1
			result = m
		} else {
			l = m + 1
		}
	}

	fmt.Println(result)
}
