package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var valsToNum map[string]int = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

func getDeck() (deck []string) {
	var (
		suit []string = []string{"S", "C", "D", "H"}
		vals []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	)
	for _, n := range vals {
		for _, s := range suit {
			deck = append(deck, n+s)
		}
	}
	return
}

func delVals(list []string, vals []string) []string {
	var (
		exist bool
		res   []string
	)

	_ = copy(res, list)
	for _, val := range list {
		for _, del := range vals {
			if val == del {
				exist = true
			}
		}
		if !exist {
			res = append(res, val)
		}
		exist = false
	}
	return res
}

func calc(hand []string, card string) []int {
	res := []int{0, 0}
	v1 := valsToNum[string(hand[0][0])]
	v2 := valsToNum[string(hand[1][0])]
	v3 := valsToNum[string(card[0])]
	if v1 == v2 && v2 == v3 {
		res[0] = 3
		res[1] = v1
	} else if v1 == v2 || v1 == v3 {
		res[0] = 2
		res[1] = v1
	} else if v2 == v3 {
		res[0] = 2
		res[1] = v2
	} else if v1 > v2 && v1 > v3 {
		res[0] = 1
		res[1] = v1
	} else if v2 > v1 && v2 > v3 {
		res[0] = 1
		res[1] = v2
	} else {
		res[0] = 1
		res[1] = v3
	}
	return res
}

func checkWin(players [][]int) bool {
	win := true
	first := players[0]
	for _, info := range players {
		if info[0] > first[0] {
			win = false
			break
		}
		if info[0] == first[0] {
			if info[1] > first[1] {
				win = false
				break
			}

		}
	}
	return win
}

func main() {
	deck := getDeck()

	var (
		n, players int
		res        []string
		//queue      []int
	)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for t := 0; t < n; t++ {
		s.Scan()
		players, _ = strconv.Atoi(s.Text())

		deck = getDeck()
		cards := make([][]string, players)
		info := make(map[string][][]int)
		res = nil
		//читаем руки игроков
		for k := 0; k < players; k++ {
			s.Scan()
			cards[k] = strings.Split(s.Text(), " ") // сохраняем руки
			deck = delVals(deck, cards[k])          //удаляем с колоды руки
		}
		for _, table := range deck {
			info[table] = make([][]int, players)
			for i, card := range cards {
				info[table][i] = calc(card, table) //считаем для игрока инфу
			}
			if checkWin(info[table]) {
				res = append(res, table)
			}
		}
		fmt.Println(len(res))

		for _, val := range res {
			fmt.Println(val)

		}

	}
}
