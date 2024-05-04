package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		n  int
		dt string
	)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		dt = s.Text()
		if _, err := time.Parse("2 1 2006", dt); err != nil {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
