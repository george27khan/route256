package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var cntFile = 0

func checkFiles(files []interface{}, hack bool) (resHack bool, cnt int) {
	for _, file := range files {
		cnt++ //локальный счетчик
		if !resHack && strings.HasSuffix(file.(string), ".hack") {
			resHack = true
		}
	}
	if resHack || hack {
		cntFile += cnt
	} else {
		cnt = 0
	}
	return
}

func check(jsmap map[string]interface{}, hack bool) {
	var resHack bool
	if _, ok := jsmap["files"]; ok {
		resHack, _ = checkFiles(jsmap["files"].([]interface{}), hack)
	}
	if _, ok := jsmap["folders"]; ok {
		for _, val := range jsmap["folders"].([]interface{}) {
			if resHack || hack {
				check(val.(map[string]interface{}), true)
			} else {
				check(val.(map[string]interface{}), false)
			}
		}
	}
	return
}

// Оповещения
func main() {
	var (
		in      *bufio.Reader
		out     *bufio.Writer
		n, t    int
		js, row []byte
	)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var jsmap map[string]interface{}

		fmt.Fscan(in, &t)
		for j := 0; j <= t; j++ {
			row, _ = in.ReadBytes('\n')
			//fmt.Println("row1 ", s.Text())
			js = append(js, row...)
		}
		json.Unmarshal(js, &jsmap)
		//fmt.Println(string(js))
		//fmt.Println(jsmap)
		check(jsmap, false) // считаем вирусы
		fmt.Println(cntFile)
		js = []byte{}
		cntFile = 0
	}
}

/*
dfdgfdsg
*/
