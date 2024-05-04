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
		n, cur, next, cnt, base int
		row                     []string
		up, down                bool
		res                     []int
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		_, _ = strconv.Atoi(s.Text())
		s.Scan()
		row = strings.Split(s.Text(), " ")
		if len(row) == 1 {
			cur, _ = strconv.Atoi(row[0])
			res = append(res, cur, 0)
		}
		for j := 0; j < len(row)-1; j++ {
			cur, _ = strconv.Atoi(row[j])
			next, _ = strconv.Atoi(row[j+1])
			if !up && !down {
				if next-cur == 1 {
					base = cur
					cnt++
					up = true
				} else if next-cur == -1 {
					base = cur
					cnt--
					down = true
				} else {
					res = append(res, cur, cnt)
				}
			} else if up {
				if next-cur == 1 {
					cnt++
				} else {
					res = append(res, base, cnt)
					up = false
					cnt = 0
				}
			} else if down {
				if next-cur == -1 {
					cnt--
				} else {
					res = append(res, base, cnt)
					down = false
					cnt = 0
				}
			}
			if j == len(row)-2 { // последний шаг
				if !up && !down {
					cur, _ = strconv.Atoi(row[j+1])
					res = append(res, cur, cnt)
				} else {
					res = append(res, base, cnt)
					cnt = 0
					up = false
					down = false
				}
			}
		}
		fmt.Println(len(res))
		for _, val := range res {
			fmt.Print(val, " ")
		}

		res = nil
	}

}
