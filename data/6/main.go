package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n, i, j, stepJ, idx int
		row                 []rune
		res                 [][]string = [][]string{{""}}
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	fmt.Println(n)
	for t := 0; t < n; t++ {
		s.Scan()
		row = []rune(s.Text())
		fmt.Println(string(row))
		for _, valRune := range row {
			if valRune == 'L' {
				stepJ--
				fmt.Println("stepJ--", stepJ)
			} else if valRune == 'R' {
				stepJ++
				fmt.Println("stepJ++", stepJ)
			} else {
				if j+stepJ < 0 {
					idx = 0
				} else if j+stepJ > len(res[i]) {
					idx = len(res) - 1
				} else {
					idx = j + stepJ
				}
				fmt.Println(res[0])
				head := append([]string(nil), res[i][:idx]...)
				tail := append([]string(nil), res[i][idx:]...)
				res[i] = append(append(head, string(valRune)), tail...)
				stepJ = 0
				j = idx
				j++
				fmt.Println("res", res)
			}

			//
		}
		fmt.Println(res)
	}

}
