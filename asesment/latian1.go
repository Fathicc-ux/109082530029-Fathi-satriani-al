package main

import "fmt"

func hitungganjil(n int) int{
	var bil int
	ganjil := 0
	for i:=1; i<=n; i++{
		fmt.Scan(&bil)
		if bil % 2 == 1{
			ganjil+= 1
		}
	}
	return ganjil
}
func main() {
	var n int
	fmt.Print("Masukan banyak bilangan: ")
	fmt.Scan(&n)

	hasil:= hitungganjil(n)
	fmt.Print("Jumlah ganjil: ", hasil)
}