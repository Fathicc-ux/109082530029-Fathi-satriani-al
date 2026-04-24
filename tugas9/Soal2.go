package main

import (
	"fmt"
	"math"
)

func seluruh(arr [256] int, n int) [256] int{
	
	 for i:= 0; i<n; i++{
		fmt.Scan(&arr[i])
	 }
	 fmt.Print("nilai seluruh: ")
	 for i:= 0; i<n; i++{
		fmt.Print(arr[i], " ")
	 }
	 fmt.Println()
	 return arr
}
	
func ganjil(arr [256] int, n int){
	fmt.Print("ganjil: ")
	for i:= 0; i<n; i++{
		if i%2 == 1{
		fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func genap(arr [256] int, n int){
	fmt.Print("genap: ")
	for i:= 0; i<n; i++{
		if i%2 == 0{
		fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func kelipatan(arr [256]int, n int) {
	var x int
	fmt.Print("x: ")
	fmt.Scan(&x)

	fmt.Print("kelipatan: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func hapus(arr [256]int, n int) ([256] int, int){
	var index int
	fmt.Print("hapus index: ")
	fmt.Scan(&index)

	for i:= index; i<n-1; i++{
		arr[i] = arr[i+1]
	}
	n --
	fmt.Print("hasil hapus: ")
	for i:= 0; i<n; i++{
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	return arr, n
}
func rata(arr [256] int, n int){
	jumlah:= 0
	for i:= 0; i < n; i++{
		jumlah+= arr[i]
	}
	rata:= jumlah / n
	fmt.Println("rata rata: ", rata)
}
func simpangan_baku(arr [256] int, n int){
	var total float64
	var jumlah int
	for i:= 0; i < n; i++{
		jumlah+= arr[i]
	}
	rata:= float64(jumlah) / float64(n)

	for i:=0; i<n; i++{
		selisih:= float64(arr[i]) - (rata)
		total+= (selisih) * (selisih)
	}
	fmt.Print("simpangan baku: ", math.Sqrt(float64(total)/ float64(n)))
	fmt.Println()
}
func frekuensi(arr [256] int, n int){
	var x, hitung int
	fmt.Print("frek: ")
	fmt.Scan(&x)
	for i:=0 ;i<n; i++{
		if arr[i] == x{
			hitung++
		}
	}
	fmt.Print("frekuensi: ", hitung)
}
func main () {
	var n int
	var arr [256] int
	fmt.Print("banyak n: ")
	fmt.Scan(&n)
	
	arr = seluruh(arr, n)
	ganjil(arr, n)
	genap(arr, n)
	kelipatan(arr, n)
	arr, n = hapus(arr, n)
	rata(arr, n)
	simpangan_baku(arr, n)
	frekuensi(arr, n)
	
}