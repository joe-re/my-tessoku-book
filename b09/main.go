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
	MAX := 1501
	t := make([][]int, MAX)
	for i := 0; i < MAX; i++ {
		t[i] = make([]int, MAX)
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
		// t[ax][ay] += 1
		// t[bx+1][ay] -= 1
		// t[ax][by+1] -= 1
		// t[bx+1][by+1] += 1
		t[ax][ay] += 1
		t[bx][ay] -= 1
		t[ax][by] -= 1
		t[bx][by] += 1
	}

	for i := 0; i < MAX; i++ {
		for j := 1; j < MAX; j++ {
			t[i][j] += t[i][j-1]
		}
	}
	for i := 1; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			t[i][j] += t[i-1][j]
		}
	}
	count := 0
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			if t[i][j] > 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
