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
	n, _ := strconv.Atoi(params[0])
	q, _ := strconv.Atoi(params[1])

	input, _ = reader.ReadString('\n')
	nums := strings.Fields(input)
	prefixSum := make([]int, n)

	// 累積和の計算
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(nums[i])
		if i == 0 {
			prefixSum[i] = num
		} else {
			prefixSum[i] = prefixSum[i-1] + num
		}
	}

	// クエリの処理
	results := make([]int, q)
	for i := 0; i < q; i++ {
		input, _ = reader.ReadString('\n')
		query := strings.Fields(input)
		l, _ := strconv.Atoi(query[0])
		r, _ := strconv.Atoi(query[1])

		if l > 1 {
			results[i] = prefixSum[r-1] - prefixSum[l-2]
		} else {
			results[i] = prefixSum[r-1]
		}
	}

	// 結果の出力
	for _, result := range results {
		fmt.Println(result)
	}
}
