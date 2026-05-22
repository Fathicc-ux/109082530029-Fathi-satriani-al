// Fathi Satriani Aljauzy_109082530029
package main

import "fmt"

const arr int = 100

type himpunan [arr]int

func gabungan(A, B himpunan, nA, nB int) {
	var gabungan himpunan
	var nG int

	for i := 0; i < nA; i++ {
		gabungan[i] = A[i]
	}
	nG = nA
	for i := 0; i < nB; i++ {
		jumlah := 0
		for j := 0; j < nG; j++ {
			if B[i] == gabungan[j] {
				jumlah++
			}
		}
		if jumlah == 0 {
			gabungan[nG] = B[i]
			nG++
		}

	}
	for i := 0; i < nG; i++ {
		fmt.Print(gabungan[i], " ")
	}
}
func main() {

	A := himpunan{11, 28, 33, 64, 95, 16, 100, 15}
	B := himpunan{3, 11, 7, 28, 33, 6}
	nA := 8
	nB := 6
	fmt.Print("Gabungan A dan B: ")
	gabungan(A, B, nA, nB)
}