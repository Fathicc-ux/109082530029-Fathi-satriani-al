//Fathi Satriani Al_109082530029
package main

import "fmt"

func fx(a int)int{
	hasil:= a * a
	return hasil
}
func gx(a int)int{
	hasil:= a -2
	return hasil
}
func hx(a int)int{
	hasil:= a + 1
	return hasil
}
func main() {
	var a, b, c int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}