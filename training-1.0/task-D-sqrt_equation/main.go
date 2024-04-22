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
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	fmt.Fprintln(out, Solution(a, b, c))
}

func Solution(a, b, c int) string {
	if c < 0 {
		return "NO SOLUTION"
	}
	if a == 0 {
		if c*c == b {
			return "MANY SOLUTIONS"
		}
		return "NO SOLUTION"
	}
	if (c*c-b)%a == 0 {
		return fmt.Sprint((c*c - b) / a)
	}
	return "NO SOLUTION"
}
