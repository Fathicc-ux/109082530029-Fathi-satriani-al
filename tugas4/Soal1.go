package main

import "fmt"

func faktorial(n int, hasil* int){
	*hasil = 1
	for i:=1; i<=n; i++{
		*hasil = *hasil * i
	}
}
func permutasi(n, r int, hasil* int) {
    var nf, fn int
    faktorial(n, &nf)
    faktorial(n-r, &fn)
	*hasil = nf/fn
}
func kombinasi(n, r int, hasil* int) {
    var nf, nr, fn int
    faktorial(n, &nf)
    faktorial(r, &nr)
    faktorial(n-r, &fn)
	*hasil = nf / (nr * fn)
}
func main () {
	var a, b, c, d int
	var per1, per2, kom1, kom2 int

	fmt.Scan(&a, &b, &c, &d)
	
	permutasi(a, c, &per1)
	kombinasi(a, c, &kom1)
	permutasi(b, d, &per2)
	kombinasi(b, d, &kom2)
	
	fmt.Println("Baris 1: ", per1, kom1)
	fmt.Println("Baris 2: ", per2, kom2)
}