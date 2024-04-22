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
	ans := Solution(a, b, c)
	fmt.Fprintln(out, ans)
}

func Solution(a, b, c int) string {
	if a+b > c && a+c > b && b+c > a {
		return "YES"
	}
	return "NO"
}
