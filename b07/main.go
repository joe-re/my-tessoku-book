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
	t, _ := strconv.Atoi(params[0])

	input, _ = reader.ReadString('\n')
	params = strings.Fields(input)
	n, _ := strconv.Atoi(params[0])
	diff := make([]int, t)

	for i := 0; i < n; i++ {
		input, _ = reader.ReadString('\n')
		params = strings.Fields(input)
		l, _ := strconv.Atoi(params[0])
		r, _ := strconv.Atoi(params[1])
		diff[l]++
		if r < t {
			diff[r]--
		}
	}

	current := 0
	for i := 0; i < t; i++ {
		current += diff[i]
		fmt.Println(current)
	}
}
