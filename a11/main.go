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

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	l := 0
	r := n

	result := -1
	for l < r {
		m := (l + r) / 2
		if a[m] == x {
			result = m
			break
		}
		if a[m] < x {
			l = m
		} else {
			r = m
		}
	}

	fmt.Println(result + 1)
}
