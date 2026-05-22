//Fathi Satriani Aljauzy_109082530029
package main

import "fmt"

const arr int = 100

type himpunan [arr] int
func valid(A himpunan, n int){
	var frekuensi [101] int
	jumlah:= 0

	for i:= 0; i<n; i++{
		frekuensi[A[i]]++
	}
	for i:=0; i<=100; i++{
		if frekuensi[i]>1{
			jumlah++
		}
	}
	if jumlah > 0{
		fmt.Println("Tidak valid")
	}else {
		fmt.Println("Valid")
	}

}
func main() {
	A:= himpunan {11, 28, 33, 64, 95, 16}
	B:= himpunan {11, 28, 33, 64, 95, 16, 100, 28, 33, 64, 95, 16}
	nA := 6
	nB := 12

	fmt.Print("Himpunan A: ")
	valid(A, nA)
	fmt.Print("Himpunan B: ")
	valid(B, nB)
}