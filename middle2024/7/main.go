package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

var (
	bestPart   string
	bestCnt    int
	sched      []int
	black, row string
	deck       map[string]int
	cashCycle  map[string]int
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

type Person struct {
	id     int
	time   int
	action string
}

type ByTime []Person

func (p ByTime) Len() int { return len(p) }

func (p ByTime) Less(i, j int) bool { return p[i].time < p[j].time }

func (p ByTime) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type ById []Person

func (p ById) Len() int { return len(p) }

func (p ById) Less(i, j int) bool { return p[i].id < p[j].id }

func (p ById) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func check(personList []Person, n int) (ok bool) {
	var prev Person
	//fmt.Println(personList)
	if personList[0].time > 1 {
		personList[0].time--
		personList[0].action = "-"
	} else {
		personList[0].action = "0"
	}
	prev = personList[0]
	for i := 1; i < len(personList); i++ {
		//fmt.Println(personList[i].time - prev.time)
		//если окно совпадает, то сдвигаем на 1
		if personList[i-1].time > n {
			return
		}
		if personList[i].time-prev.time == 1 {
			personList[i].action = "0"
			prev = personList[i]
			continue
		}
		if personList[i].time == prev.time {
			personList[i].time++
			if personList[i].time > n {
				return
			}
			personList[i].action = "+"
			prev = personList[i]

		} else if personList[i].time-prev.time > 1 {
			personList[i].time--
			personList[i].action = "-"
			prev = personList[i]
		} else {
			return
		}
	}
	return true
}

// Оповещения
func main() {
	//defer un(trace("calc"))
	var (
		in           *bufio.Reader
		out          *bufio.Writer
		t, n, m, seq int
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		personList := make([]Person, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &personList[j].time)
			personList[j].id = seq
			seq++
		}
		sort.Sort(ByTime(personList))
		if ok := check(personList, n); ok {
			//fmt.Println(personList)

			sort.Sort(ById(personList))
			//fmt.Println(personList)
			for _, person := range personList {
				fmt.Print(person.action)
			}
			fmt.Println()
		} else {
			fmt.Println("x")
		}
		seq = 0
		//fmt.Println(blue, red, white, black)
	}
}
