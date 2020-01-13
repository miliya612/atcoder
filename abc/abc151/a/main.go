package main

import (
	"fmt"
	"strings"
)

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	var s string
	fmt.Scan(&s)
	i := strings.Index(alpha, s)+1
	fmt.Println(string([]rune(alpha)[i]))
}

