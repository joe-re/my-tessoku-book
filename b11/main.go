package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	sort.Ints(a)

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	x := make([]int, q)

	for i := 0; i < q; i++ {
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
	}

	results := make([]int, q)

	for i := 0; i < q; i++ {
		target := x[i]
		l := 0
		r := n

		result := 0

		for l < r {
			m := (l + r) / 2
			if a[m] < target {
				result += m - l + 1
				l = m + 1
			} else {
				r = m
			}
		}

		results[i] = result
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
