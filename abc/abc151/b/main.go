package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	// 科目数
	n := nextInt()
	// 満点
	k := nextInt()
	// 目標平均点
	m := nextInt()

	points := []int{}
	total := 0
	for i:=0; i<n-1; i++ {
		p := nextInt()
		points = append(points, p)
		total+=p
	}

	rest := (n*m) - total
	if rest > k {
		fmt.Println("-1")
	} else if rest < 0 {
		fmt.Println("0")
	} else {
		fmt.Println(rest)
	}
}