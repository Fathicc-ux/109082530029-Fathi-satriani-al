package main

import "fmt"

func array(NMAX *[127] rune, n int){
	for i:=0; i<n; i++{
		var kata string
		fmt.Scan(&kata)
		NMAX [i] =  rune(kata[0])
	}
}
func revrese(NMAX *[127] rune, n int){
	var a [127] rune
	for i:=0; i<n; i++{
		a[i] = NMAX[n-1-i]
	}
	for i:= 0; i<n; i++{
		NMAX[i] = a[i] 
	}
}
func palindrom(NMAX *[127] rune, n int) bool{
	var a [127] rune
	for i:=0; i<n; i++{
		a[i] = (NMAX[n-1-i])
	}
	for i:= 0; i<n; i++{
		if NMAX[i] != a[i]{
			return false
		}
	}
	return  true
}	
func main() {
	var NMAX [127] rune
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)

	array(&NMAX, n)
	fmt.Println("ISI ARRAY:")
for i := 0; i < n; i++ {
	fmt.Print(string(NMAX[i]), " ")
}
fmt.Println()
revrese(&NMAX, n)
	if palindrom(&NMAX, n){
		fmt.Println("Palindrom")
	}else{
		fmt.Print("bukan")
	}
	
}