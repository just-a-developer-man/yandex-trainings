package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tCond, tRoom, ans int
	var condMode string

	fmt.Fscan(in, &tRoom, &tCond, &condMode)
	ans = Solution(tRoom, tCond, condMode)
	fmt.Fprintln(out, ans)
}

func Solution(tRoom, tCond int, condMode string) int {
	switch condMode {
	case "heat":
		if tRoom < tCond {
			return tCond
		}
		return tRoom
	case "freeze":
		if tRoom > tCond {
			return tCond
		}
		return tRoom
	case "auto":
		return tCond
	}
	return tRoom
}
