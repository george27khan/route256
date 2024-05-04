package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n, m, val int
		setting   []string
		low       int = 15
		high      int = 30
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		m, _ = strconv.Atoi(s.Text())
		for j := 0; j < m; j++ {
			s.Scan()
			setting = strings.Split(s.Text(), " ")
			val, _ = strconv.Atoi(setting[1])
			if setting[0] == ">=" && val > low {
				low = val
			} else if setting[0] == "<=" && val < high {
				high = val
			}
			if low > high {
				fmt.Println("-1")
			} else {
				fmt.Println(low)
			}
		}
		low = 15
		high = 30
	}

}
