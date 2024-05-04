package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		n, cnt int
		res    []string
		//queue      []int
	)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for t := 0; t < n; t++ {
		s.Scan()
		cnt, _ = strconv.Atoi(s.Text())
		//читаем руки игроков
		for k := 0; k < cnt; k++ {
			s.Scan()
			cards[k] = strings.Split(s.Text(), " ") // сохраняем руки
			deck = delVals(deck, cards[k])          //удаляем с колоды руки
		}

	}
}
