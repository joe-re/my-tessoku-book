package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findIndex(x []int, k int) int {
	left := 0
	right := len(x) - 1
	for left <= right {
		mid := (left + right) / 2
		if x[mid] == k {
			return mid
		} else if x[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	a2 := make([]int, n)
	copy(a2, a)
	sort.Ints(a2)
	x := make([]int, 0)

	prev := -1
	for i := 0; i < n; i++ {
		if prev != a2[i] {
			x = append(x, a2[i])
			prev = a2[i]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(findIndex(x, a[i]) + 1)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}
