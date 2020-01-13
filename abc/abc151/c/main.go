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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	// 問題数
	_ = nextInt()
	// 提出回数
	m := nextInt()

	was := map[int]int{}
	acs := map[int]bool{}
	for i := 0; i < m; i++ {
		j := nextInt()
		res := nextString()
		if _,ok := acs[j]; ok {
			continue
		}
		if res  == "WA" {
			was[j]++
		} else {
			acs[j] = true
		}
	}

	finAcs := len(acs)
	finWas := 0

	for j,_ := range acs {
		if r,ok := was[j]; ok {
			finWas += r
		}
	}

	fmt.Printf("%d %d\n" ,finAcs, finWas)
}
