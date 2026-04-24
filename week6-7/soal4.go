// Fathi satriani Al_109082530029
package main

import "fmt"

func input(t *[256]int) int {
	var gol int
	n := 0
	for {
		fmt.Scan(&gol)
		if gol < 0 {
			return n
		}
		t[n] = gol
		n++
	}
}
func rata(t [256]int, n int) int {
	total := 0

	for i := 0; i < n; i++ {
		total += t[i]
	}

	return (total) / (n)
}

func main() {
	var gol1, gol2 [256]int
	var data1, data2 int
	var r1, r2 int

	data1 = input(&gol1)
	data2 = input(&gol2)

	r1 = rata(gol1, data1)
	r2 = rata(gol2, data2)

	fmt.Println(r1)
	fmt.Println(r2)
}
