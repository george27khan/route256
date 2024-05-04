package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	bestPart         string
	bestCnt          int
	blue, red, white []string
	black, row       string
	deck             map[string]int
	blueSet          map[string]int
	redSet           map[string]int
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func c(deck []string, str string) (cnt int) {
	//defer un(trace("c"))

	for _, val := range deck { //
		if strings.Contains(val, str) {
			cnt++
		}
	}
	return
}
func subset_string(string_val string) (res []string) {
	//defer un(trace("subset_string"))
	size := len(string_val) //denotes length of input string
	step := 1
	cash := make(map[string]int)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j+step > size {
				break
			}
			str := string_val[j : j+step]
			if _, ok := cash[str]; !ok {
				res = append(res, str)
				cash[str] = 1
			}
		}
		step++
	}
	return res //[1 : len(res)-1]
}

func calc() {
	//defer un(trace("calc"))

	for _, val := range blue {
		parts := subset_string(val)
		//fmt.Println(parts)
		for _, part := range parts {
			if _, ok := deck[part]; ok || strings.Contains(black, part) { // если фраза из колоды или часть черного слова, то пропускаем фразу
				continue
			}
			if _, ok := blueSet[part]; ok {
				blueSet[part]++
			} else {
				blueSet[part] = 1
			}
		}
	}

	for _, val := range red {
		parts := subset_string(val)
		for _, part := range parts {
			if _, ok := deck[part]; ok || strings.Contains(black, part) { // если фраза из колоды или часть черного слова, то пропускаем фразу
				continue
			}
			if _, ok := redSet[part]; ok {
				redSet[part]++
			} else {
				redSet[part] = 1
			}
		}
	}
	for key, val := range blueSet {
		diff := val - redSet[key]
		if diff > bestCnt {
			bestCnt = diff
			bestPart = key
		}
	}
	//проверка черного слова
	//проверка словаря
	//подсчет метрики
}

// Оповещения
func main() {
	//fmt.Println(subset_string("ccbbaa"))
	//return
	var (
		in            *bufio.Reader
		out           *bufio.Writer
		t, n, b, r, f int
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &b, &r, &f)
		blue = make([]string, b)
		red = make([]string, r)
		white = make([]string, n-b-r-1)
		deck = make(map[string]int)
		blueSet = make(map[string]int)
		redSet = make(map[string]int)

		for j := 0; j < n; j++ {
			if j < b {
				fmt.Fscan(in, &row)
				blue[j] = row
			} else if j >= b && j < b+r {
				fmt.Fscan(in, &row)
				red[j-b] = row
			} else if j == f-1 {
				fmt.Fscan(in, &row)
				black = row
			} else {
				fmt.Fscan(in, &row)
			}
			deck[row] = 1
		}
		calc()
		if bestCnt > 0 {
			fmt.Println(bestPart, bestCnt)
			//fmt.Println(blueSet)
			//fmt.Println(redSet)

		} else {
			fmt.Println("dfsfgh", bestCnt)
		}
		bestPart = ""
		bestCnt = 0
		//fmt.Println(blue, red, white, black)
	}
}
