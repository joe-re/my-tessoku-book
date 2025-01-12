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
	k, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	res := 0
	for left := 0; left < n && a[left] <= k; left++ {
		sum := a[left]
		right := left
		for right+1 < n && sum+a[right+1] <= k {
			sum += a[right+1]
			right++
		}
		// fmt.Println(right, left, sum)
		res += right - left + 1
	}

	fmt.Println(res)
}
