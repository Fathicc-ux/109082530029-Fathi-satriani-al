package main

import "fmt"

	
func denom(uang int){
	k10 := uang/10000
	sis := uang%10000
	k2 := sis/2000
	sis = sis % 2000
	k1 := sis/1000
	fmt.Println(k10, "lembar 10000")
	fmt.Println(k2, "lembar 2000")
	fmt.Println(k1, "lembar 1000")
	}
func main() {
	var uang int
	fmt.Scan(&uang)

	denom(uang)
}