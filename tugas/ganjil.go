package main

import "fmt"

func ganjil(n int){
	if n == 0{
		 return
	}
	// digit:= n%10
	// ganjil(n/10)
	ganjil(n-1)
	if n%2 != 0{
		fmt.Print(n, " ")
	}
}
func main() {
	var n int
	fmt.Scan(&n)

	ganjil(n)
}