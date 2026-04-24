//Fathi Satriani Al_109082530029
package main

import "fmt"

func faktorial(n int)int{
		hasil:= 1
		for i:=1; i<=n; i++{
			hasil = hasil * i
		}
	return hasil
}
func permutasi(n, r int)int{
	hasil := faktorial(n)/faktorial(n-r)
	return hasil
}
func main() {
	var x, y int

	fmt.Scan(&x, &y)

	fmt.Print(faktorial(x), faktorial(y), " ")
	if x >= y{
		fmt.Print(permutasi(x, y))
	}else{
		fmt.Print(permutasi(y, x))
	}
}