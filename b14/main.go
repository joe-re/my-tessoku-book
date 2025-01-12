package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func enumerate(ary []int) []int {
	result := make([]int, 1<<len(ary))
	for i := 0; i < 1<<len(ary); i++ {
		sum := 0
		for j := 0; j < len(ary); j++ {
			if i&(1<<j) != 0 {
				sum += ary[j]
			}
		}
		result[i] = sum
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	ary := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		ary[i], _ = strconv.Atoi(scanner.Text())
	}

	mid := n / 2
	p := enumerate(ary[:mid])
	q := enumerate(ary[mid:])

	sort.Ints(p)
	sort.Ints(q)

	for _, v := range p {
		left := 0
		right := len(q) - 1
		for left <= right {
			mid := (left + right) / 2
			if v+q[mid] == k {
				fmt.Println("Yes")
				return
			} else if v+q[mid] < k {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	fmt.Println("No")
}
