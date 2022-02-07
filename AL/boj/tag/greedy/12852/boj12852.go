package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var X int
	fmt.Fscanf(reader, "%d\n", &X)

	dp := make([]int, X+1)
	for i := 2; i < X+1; i++ {
		dp[i] = 100001
	}

	parent := make([]int, X+1)
	for i := 2; i < X+1; i++ {
		parent[i] = i
	}

	for i := 2; i < X+1; i++ {
		if i%2 == 0 {
			if dp[i] > dp[i/2]+1 {
				dp[i] = dp[i/2] + 1
				parent[i] = i / 2
			}
		}

		if i%3 == 0 {
			if dp[i] > dp[i/3]+1 {
				dp[i] = dp[i/3] + 1
				parent[i] = i / 3
			}
		}

		if dp[i] > dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
			parent[i] = i - 1
		}
	}

	fmt.Fprintln(writer, dp[X])
	K := X
	for K != 0 {
		fmt.Fprintf(writer, "%d ", K)
		K = parent[K]
	}
}
