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
	n, m, seq int
	friends   []Friend
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

type Friend struct {
	id      int
	cards   int
	newCard int
}

type ByCards []Friend

func (p ByCards) Len() int { return len(p) }

func (p ByCards) Less(i, j int) bool { return p[i].cards > p[j].cards }

func (p ByCards) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type ById []Friend

func (p ById) Len() int { return len(p) }

func (p ById) Less(i, j int) bool { return p[i].id < p[j].id }

func (p ById) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func check(deck []int) (ok bool) {
	//var prev Person
	deckLen := len(deck)
	for i, friend := range friends {
		if friend.cards < deckLen-i && deck[deckLen-1-i] == 0 {
			deck[deckLen-1-i] = 1
			friends[i].newCard = deckLen - i
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
		in  *bufio.Reader
		out *bufio.Writer
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &m, &n)
	friends = make([]Friend, m)
	deck := make([]int, n)
	for j := 0; j < m; j++ {
		fmt.Fscan(in, &friends[j].cards)
		friends[j].id = seq
		seq++
	}
	sort.Sort(ByCards(friends))
	if ok := check(deck); ok {
		//fmt.Println(personList)
		sort.Sort(ById(friends))
		//fmt.Println(personList)
		for _, friend := range friends {
			fmt.Print(friend.newCard, " ")
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}

	//fmt.Println(blue, red, white, black)
}
