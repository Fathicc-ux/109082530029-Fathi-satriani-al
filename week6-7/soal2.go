// Fathi satriani Al_109082530029
package main

import "fmt"

type persegi struct{
	panjang, lebar int
	warna string
}
type geometri struct{
	luas, keliling int
}
func input(in *persegi){
	fmt.Scan(&in.panjang, &in.lebar)
	fmt.Scan(&in.warna)
}
func hitung(hit *geometri, in persegi){
	hit.luas = in.panjang * in.lebar
	hit.keliling = 2 * (in.panjang + in.lebar)
	fmt.Println("Luas: ", hit.luas, "warna: ", in.warna)
	fmt.Print("keliling: ", hit.keliling, " warna: ", in.warna)
}
func main() {
	var in persegi
	var hit geometri
	input(&in)
	hitung(&hit, in)
}