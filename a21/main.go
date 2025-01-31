package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力からデータを読み込むためのスキャナを作成
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Nを読み込む
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// PとAの配列を作成
	P := make([]int, N+1)
	A := make([]int, N+1)

	// PとAの値を読み込む
	for i := 1; i <= N; i++ {
		scanner.Scan()
		P[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}

	// dpテーブルを作成
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}

	// 初期状態を設定
	dp[1][N] = 0

	// 動的計画法のループ
	for LEN := N - 2; LEN >= 0; LEN-- {
		for l := 1; l <= N-LEN; l++ {
			r := l + LEN

			// score1の計算
			score1 := 0
			if l >= 2 && l <= P[l-1] && P[l-1] <= r {
				score1 = A[l-1]
			}

			// score2の計算
			score2 := 0
			if r <= N-1 && l <= P[r+1] && P[r+1] <= r {
				score2 = A[r+1]
			}

			// dp[l][r]の計算
			if l == 1 {
				dp[l][r] = dp[l][r+1] + score2
			} else if r == N {
				dp[l][r] = dp[l-1][r] + score1
			} else {
				dp[l][r] = max(dp[l-1][r]+score1, dp[l][r+1]+score2)
			}
		}
	}

	answer := 0
	for i := 1; i <= N; i++ {
		answer = max(answer, dp[i][i])
	}
	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
