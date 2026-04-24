//Fathi satriani Al_109082530029
package main

import "fmt"

type nasabah struct{
	id_nasabah, norek, nama, bank string
}
type tabnasabah [2022]nasabah

func addNasabah(t *tabnasabah, n *int, data nasabah){
	if *n >= 2022{
		fmt.Println("penuh")
		return
	}
	t[*n] = data
	*n = *n + 1
}
func cetak(t tabnasabah, n int, x string){
	for i:= 0; i<n; i++{
		if t[i].bank == x{
			fmt.Println(t[i])
		}
	}
}

func main() {
	var t tabnasabah
	var n, input int
	var d nasabah
	var caribank string
	fmt.Print("jumlah input: ")
fmt.Scan(&input)
for i:=1; i<=input; i++{

	fmt.Scan(&d.id_nasabah, &d.nama, &d.bank, &d.norek)
	fmt.Println("--------")

	addNasabah(&t, &n, d)
}
	fmt.Print("Cari Bank: ")
	fmt.Scan(&caribank)
cetak(t, n, caribank)
}
