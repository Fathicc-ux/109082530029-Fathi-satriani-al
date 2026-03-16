package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c ,d float64)float64{
	hasil:= math.Sqrt((a-c)*(a-c) + (b - d)*(b-d))
	return hasil
}
func dalam(cx, cy, r, x, y float64)bool{
	if jarak(x, y, cx, cy) <= r{
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

	daling1:= dalam(float64(a1), float64(b1), float64(r1), float64(x), float64(y))
	daling2:= dalam(float64(a2), float64(b2), float64(r2), float64(x), float64(y))

	if daling1 && daling2{
		fmt.Println("Titik di dalam lingkaran 1 & 2")
	}else if daling1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if daling2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 & 2")
	}
}