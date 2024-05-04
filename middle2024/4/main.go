package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	id    int
	time  int
	place int
}

type ByTime []Person

func (p ByTime) Len() int { return len(p) }

func (p ByTime) Less(i, j int) bool { return p[i].time < p[j].time }

func (p ByTime) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type ById []Person

func (p ById) Len() int { return len(p) }

func (p ById) Less(i, j int) bool { return p[i].id < p[j].id }

func (p ById) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Оповещения
func main() {
	var (
		in            *bufio.Reader
		out           *bufio.Writer
		seq           = 1
		cntPlace      = 1
		prevVal, n, t int
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		personList := make([]Person, t)
		for j := 0; j < t; j++ {
			fmt.Fscan(in, &personList[j].time)
			personList[j].id = seq
			seq++
		}
		seq = 1
		sort.Sort(ByTime(personList))
		for j, person := range personList {
			if person.time > prevVal && person.time-prevVal > 1 {
				cntPlace = seq
			}
			personList[j].place = cntPlace
			prevVal = person.time
			seq++
		}
		sort.Sort(ById(personList))
		for _, person := range personList {
			fmt.Print(person.place, " ")
		}
		fmt.Println()
		seq, prevVal, cntPlace = 0, 0, 1
	}
}
