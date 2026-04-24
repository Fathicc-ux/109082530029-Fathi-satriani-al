package main

import "fmt"

func hitungluas(panjang int, lebar float64) float64{
		return float64(panjang) * float64(lebar)
}
func main() {
	var p int
	var l float64

	fmt.Scan(&p, &l)
	hitung:= hitungluas(p, l)
	fmt.Print(hitung)
}