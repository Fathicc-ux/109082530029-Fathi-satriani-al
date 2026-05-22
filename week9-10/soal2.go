//Fathi Satriani Aljauzy_109082530029
package main

import "fmt"

const arr int = 100

type himpunan [arr] int

func irisan(A, B himpunan, nA, nB int){
	for i:=0; i<nA; i++{
		for j:=0; j<nB; j++{
			if A[i] == B[j]{
				fmt.Print(A[i], " ")
			}
		}
	}
}
func main(){
	A:= himpunan {11, 28, 33, 64, 95, 16, 100, 15}
	B:= himpunan {3, 11, 7, 28, 33, 6}
	nA:=8
	nB:=6
	
	fmt.Print("irisan A B: ")
	irisan(A, B, nA, nB)
}