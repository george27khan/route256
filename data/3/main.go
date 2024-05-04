package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var (
		n    int
		str  []rune
		word []rune
		res  []rune
	)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		str = []rune(s.Text())
		for j, char := range str {
			word = append(word, char)
			if len(word) == 4 {
				if matched, _ := regexp.MatchString(`\D\d\D\D`, string(word)); matched {
					//fmt.Println(string(word))
					if res != nil {
						res = append(res, ' ')
					}
					res = append(res, word...)
					word = nil
				} else if j+1 == len(str) {
					res = []rune("-")
					break
				}
			} else if len(word) == 5 {
				if matched, _ := regexp.MatchString(`\D\d\d\D\D`, string(word)); matched {
					if res != nil {
						res = append(res, ' ')
					}
					res = append(res, word...)
					word = nil
				} else {
					res = []rune("-")
					break
				}

			} else if j+1 == len(str) && len(word) < 4 {
				res = []rune("-")
				break
			}
		}
		word = nil
		fmt.Println(string(res))
		res = nil
	}
}
