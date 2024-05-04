// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//
// )
//
// var (
//
//	users map[int]int
//
// )
//
//	func sendMess(mesId *int, userId *int) {
//		if *userId == 0 {
//			for key, _ := range users {
//				users[key] = *mesId
//			}
//		} else {
//			users[*userId] = *mesId
//		}
//	}
//
//	func showMess(userId *int) {
//		fmt.Println(users[*userId])
//	}
//
// // Наклейки
//
//	func main() {
//		var in *bufio.Reader
//		var out *bufio.Writer
//		in = bufio.NewReader(os.Stdin)
//		out = bufio.NewWriter(os.Stdout)
//		defer out.Flush()
//		var (
//			n, q, oper, userId, mesSeq int
//		)
//		fmt.Fscan(in, &n, &q)
//		users = make(map[int]int)
//		for i := 1; i <= n; i++ {
//			users[i] = 0
//		}
//
//		for i := 0; i < q; i++ {
//			fmt.Fscan(in, &oper, &userId)
//			if oper == 1 { //send mess
//				mesSeq += 1
//				sendMess(&mesSeq, &userId)
//			}
//			if oper == 2 { //send mess
//				showMess(&userId)
//			}
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	users     []int
	globalMes = 0
)

func sendMess(mesId int, userId int) {
	if userId == 0 {
		globalMes = mesId
	} else {
		users[userId-1] = mesId
	}
}

func showMess(userId int) {
	lastMes := users[userId-1]
	if globalMes < lastMes {
		fmt.Println(users[userId-1])
	} else {
		fmt.Println(globalMes)
	}
}

// Оповещения
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var (
		n, q, oper, userId, mesSeq int
	)
	fmt.Fscan(in, &n, &q)
	users = make([]int, n)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &oper, &userId)
		if oper == 1 { //send mess
			mesSeq += 1
			sendMess(mesSeq, userId)
		}
		if oper == 2 { //send mess
			showMess(userId)
		}
	}
}
