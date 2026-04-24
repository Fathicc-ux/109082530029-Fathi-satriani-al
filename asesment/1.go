package main

import "fmt"

func rata(n int) int{
	var angka int
	total:= 0
	for i:=1; i<=n; i++{
	fmt.Scan(&angka)
	total = total + angka
	} 
	rata:= total/n

	return rata
}
func main() {
	var n  int
	fmt.Print("banyak bilangan: ")
	fmt.Scan(&n)
	hasil:= rata(n)
	fmt.Print(hasil)
}