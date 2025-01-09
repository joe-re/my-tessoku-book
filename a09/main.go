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
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	t := make([][]int, h+2)
	for i := 0; i < h+2; i++ {
		t[i] = make([]int, w+2)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		ax, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		ay, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		bx, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		by, _ := strconv.Atoi(scanner.Text())

		// 座標を 0-based に調整
		ax--
		ay--
		bx--
		by--

		t[ax][ay] += 1
		t[bx+1][ay] -= 1
		t[ax][by+1] -= 1
		t[bx+1][by+1] += 1
	}

	for i := 0; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			t[i][j] += t[i][j-1]
		}
	}

	for i := 1; i < h+1; i++ {
		for j := 0; j < w+1; j++ {
			t[i][j] += t[i-1][j]
		}
	}

	for i := 0; i < h; i++ {
		strs := make([]string, w)
		for j := 0; j < w; j++ {
			strs[j] = strconv.Itoa(t[i][j])
		}
		fmt.Println(strings.Join(strs, " "))
	}
}
