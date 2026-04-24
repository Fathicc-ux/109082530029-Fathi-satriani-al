package main

import "fmt"

func array(NMAX *[127] rune, n *int){
	var kata string
	i:=0
	fmt.Print("isi: ")
	for {
		fmt.Scan(&kata)
		if kata == "."{
			break
		}
		(*NMAX) [i] =  rune(kata[0])
		i++
	}
	*n = i
}
func balikan(NMAX *[127] rune, n int){
	fmt.Print("balik: ")
	for i:=n-1; i>=0; i--{
		 fmt.Print(string((*NMAX)[i]), " ")
	}
}
func palindrom(NMAX *[127] rune, n int) bool{
	for i:=0; i<n/2; i++{
		if(*NMAX)[i] != (*NMAX)[n-1-i]{
			return false
		}
	}
	return  true
}	
func main() {
	var NMAX [127] rune
	var n int

	array(&NMAX, &n)
	fmt.Print("cetak array: ")
for i := 0; i < n; i++ {
	fmt.Print(string(NMAX[i]), " ")
}
fmt.Println()
	 balikan(&NMAX, n)
fmt.Println()
	if palindrom(&NMAX, n){
		fmt.Println("Palindrom")
	}else{
		fmt.Print("bukan")
	}
}