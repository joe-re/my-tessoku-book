package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	params := strings.Fields(input)
	d, _ := strconv.Atoi(params[0])

	input, _ = reader.ReadString('\n')
	params = strings.Fields(input)
	n, _ := strconv.Atoi(params[0])

	diff := make([]int, d)
	for i := 0; i < n; i++ {
		input, _ = reader.ReadString('\n')
		nums := strings.Fields(input)
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		diff[l-1]++
		if r < d {
			// 最終日は減少がない
			diff[r]--
		}
	}

	current := 0
	for i := 0; i < d; i++ {
		current += diff[i]
		fmt.Println(current)
	}
}
