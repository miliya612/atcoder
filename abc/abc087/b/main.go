package main

import "fmt"

func main() {
	// a: 500, b: 100, c: 50
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	var cnt int
	for ax := a; ax >= 0; ax-- {
		for bx := b; bx >= 0; bx-- {
			for cx := c; cx >= 0; cx-- {
				if (500*ax + 100*bx + 50*cx) == x {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
