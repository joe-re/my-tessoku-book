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
	h, _ := strconv.Atoi(params[0])
	w, _ := strconv.Atoi(params[1])

	twoDimSum := make([][]int, h)

	for i := 0; i < h; i++ {
		twoDimSum[i] = make([]int, w)
		input, _ = reader.ReadString('\n')
		nums := strings.Fields(input)
		wnum := 0
		for j := 0; j < w; j++ {
			cnum, _ := strconv.Atoi(nums[j])
			wnum += cnum
			twoDimSum[i][j] = wnum
		}
	}

	for i := 0; i < w; i++ {
		hnum := 0
		for j := 0; j < h; j++ {
			hnum += twoDimSum[j][i]
			twoDimSum[j][i] = hnum
		}
	}

	input, _ = reader.ReadString('\n')
	params = strings.Fields(input)
	q, _ := strconv.Atoi(params[0])
	results := make([]int, q)

	for i := 0; i < q; i++ {
		input, _ = reader.ReadString('\n')
		nums := strings.Fields(input)
		_ax, _ := strconv.Atoi(nums[0])
		ax := _ax - 1
		_ay, _ := strconv.Atoi(nums[1])
		ay := _ay - 1
		_bx, _ := strconv.Atoi(nums[2])
		bx := _bx - 1
		_by, _ := strconv.Atoi(nums[3])
		by := _by - 1
		rectEnd := twoDimSum[bx][by]
		outRectLeftTop := 0
		if ax > 0 && ay > 0 {
			outRectLeftTop = twoDimSum[ax-1][ay-1]
		}
		outRectRightTop := 0
		if ay > 0 {
			outRectRightTop = twoDimSum[bx][ay-1]
		}
		outRectLeftBottom := 0
		if ax > 0 {
			outRectLeftBottom = twoDimSum[ax-1][by]
		}
		results[i] = rectEnd + outRectLeftTop - outRectRightTop - outRectLeftBottom
	}

	for _, result := range results {
		fmt.Println(result)
	}

	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		fmt.Println(twoDimSum[i][j])
	// 	}
	// }
}
