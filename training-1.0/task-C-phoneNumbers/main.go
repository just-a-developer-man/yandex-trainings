package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var number string
	var phoneBook [3]string

	fmt.Fscan(in, &number, &phoneBook[0], &phoneBook[1], &phoneBook[2])
	number = toStandardNumber(number)
	phoneBook[0] = toStandardNumber(phoneBook[0])
	phoneBook[1] = toStandardNumber(phoneBook[1])
	phoneBook[2] = toStandardNumber(phoneBook[2])

	for _, phoneBookNumber := range phoneBook {
		fmt.Fprintln(out, Solution(number, phoneBookNumber))
	}
}

func Solution(number string, phoneBookNumber string) string {
	if number == phoneBookNumber {
		return "YES"
	}
	return "NO"
}

func toStandardNumber(number string) string {
	number = strings.ReplaceAll(number, "(", "")
	number = strings.ReplaceAll(number, ")", "")
	number = strings.ReplaceAll(number, "+7", "8")
	number = strings.ReplaceAll(number, "-", "")
	if len(number) == 7 {
		number = "8495" + number
	}
	return number
}
