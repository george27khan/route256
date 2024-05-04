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
		n, n1, n2, n3, n4 int
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		text := s.Text()
		valList := strings.Split(text, " ")
		for _, val := range valList {
			if val == "1" {
				n1++
			} else if val == "2" {
				n2++
			} else if val == "3" {
				n3++
			} else if val == "4" {
				n4++
			}
		}
		if n1 == 4 && n2 == 3 && n3 == 2 && n4 == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		n1 = 0
		n2 = 0
		n3 = 0
		n4 = 0
	}
}
