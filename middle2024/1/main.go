package main

import (
	"bufio"
	"fmt"
	"os"
)

// Сумматор
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
