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
fmt.Scan(&input)
for i:=0; i<input; i++{
	fmt.Print("data ke: ", i)
	fmt.Scanln(&d.id_nasabah)
	fmt.Scanln(&d.nama)
	fmt.Scanln(&d.bank)
	fmt.Scan(&d.norek)

	addNasabah(&t, &n, d)
}
cetak(t, n, "MANDIRI")
}
