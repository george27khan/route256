package main

import (
	"bufio"
	"fmt"
	"os"
)

// Наклейки
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		str, part     string
		strRune       []rune
		start, end, n int
	)
	fmt.Fscan(in, &str)
	fmt.Fscan(in, &n)

	strRune = []rune(str)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &start, &end, &part)
		copy(strRune[start-1:end], []rune(part))
	}
	fmt.Println(string(strRune))
}
