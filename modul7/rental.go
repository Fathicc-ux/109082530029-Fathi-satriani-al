package main

import "fmt"

type waktu struct{
	jam, menit, detik, denda int
	nama, naga string
}
func main() {
	var ren waktu
	var
	fmt.Print("Nama peneywa: ")
	fmt.Scan(&ren.nama)
	fmt.Print("Jenis naga: ")
	fmt.Scan(&ren.naga)

}