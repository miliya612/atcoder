package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	a, _ := strconv.Atoi(string(s[0]))
	b, _ := strconv.Atoi(string(s[1]))
	c, _ := strconv.Atoi(string(s[2]))
	var cnt int
	if a == 1 {
		cnt++
	}
	if b == 1 {
		cnt++
	}
	if c == 1 {
		cnt++
	}
	fmt.Println(cnt)
}
