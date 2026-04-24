package main

import (
	"fmt"
	"math"
)
type titik struct {
	x, y float64
}

type lingkaran struct {
	x, y, r float64
}
func jarak(p, q titik) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}
func dalam(c lingkaran, p titik) bool {
	if jarak(titik{c.x, c.y}, p) <= c.r {
		return true
	}
	return false
}
func main () {
	var a1, b1, r1 int
	var a2, b2, r2 int
	var x, y int

	fmt.Scanln(&a1, &b1, &r1)
	fmt.Scanln(&a2, &b2, &r2)
	fmt.Scanln(&x, &y)

	daling1 := dalam(
		lingkaran{float64(a1), float64(b1), float64(r1)},
		titik{float64(x), float64(y)},
	)

	daling2 := dalam(
		lingkaran{float64(a2), float64(b2), float64(r2)},
		titik{float64(x), float64(y)},
	)

	if daling1 && daling2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if daling1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if daling2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}