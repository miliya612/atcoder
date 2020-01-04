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
	var n int
	fmt.Scan(&n)
	sc.Split(bufio.ScanWords)

	var cnt int
	nums := []int{}
	res := true

	for i := 0; i < n; i++ {
		v := nextInt()
		if v%2 == 1 {
			res = false
			break
		}
		nums = append(nums, v)
	}

	for res {
		cnt++
		for i, v := range nums {
			if (v/2)%2 == 1 {
				res = false
				break
			}
			nums[i] = v / 2
		}
	}

	fmt.Println(cnt)
}
