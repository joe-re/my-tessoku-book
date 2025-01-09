package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)

	tolerance := 0.001
	l, r := 0.0, float64(n)+1.0

	for l < r {
		m := (l + r) / 2
		diff := n - (m*m*m + m)
		if math.Abs(diff) < tolerance {
			fmt.Printf("%.6f\n", m)
			break
		}
		if diff > 0.0 {
			l = m
		} else {
			r = m
		}
	}
}
