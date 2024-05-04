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
		n     int
		row   []string
		res   string
		queue []int
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for t := 0; t < n; t++ {
		s.Scan()
		printLen, _ := strconv.Atoi(s.Text())
		counter := make([]int, printLen)

		s.Scan()
		row = strings.Split(s.Text(), ",")
		for _, val := range row {
			if strings.Count(val, "-") == 1 {
				listRange := strings.Split(val, "-")
				listNumStart, _ := strconv.Atoi(listRange[0])
				listNumEnd, _ := strconv.Atoi(listRange[1])
				for i := listNumStart; i <= listNumEnd; i++ {
					counter[i-1] = 1
				}
			} else {
				listNum, _ := strconv.Atoi(val)
				counter[listNum-1] = 1
			}
		}
		for i, val := range counter {
			if val == 0 {
				queue = append(queue, i+1)
			} else {
				if len(queue) == 1 {
					res += strconv.Itoa(queue[0]) + ","
					queue = nil
				} else if len(queue) > 1 {
					res += strconv.Itoa(queue[0]) + "-" + strconv.Itoa(queue[len(queue)-1]) + ","
					queue = nil
				}
			}
			if val == 0 && i == len(counter)-1 {
				if len(queue) == 1 {
					res += strconv.Itoa(queue[0])
					queue = nil
				} else if len(queue) > 1 {
					res += strconv.Itoa(queue[0]) + "-" + strconv.Itoa(queue[len(queue)-1])
					queue = nil
				}
			}
		}
		fmt.Println(strings.TrimRight(res, ","))
		res = ""
	}
}
