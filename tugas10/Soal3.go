package main

import "fmt"

type arrbalita [100] float64

func hitung(arrberat *arrbalita, n int){
	for i:=0; i<n; i++{
		fmt.Print("Berat balita ke-", i+1," : ")
		fmt.Scan(&arrberat[i])
	}
	min := arrberat[0]
	max := arrberat[0]
	for i:=1; i<n; i++{
		if arrberat[i] < min{
			min = arrberat[i]
		}else if arrberat[i] > max{
			max = arrberat[i]
		}
	}
	fmt.Println("berat balita terkecil: ", min, "kg")
	fmt.Println("berat balita terbesar: ", max, "kg")
}
func rata(arrberat *arrbalita, n int)float64{
	total:= 0.0
	for i:=0 ;i<n; i++{
		total = total + arrberat[i]
	}
	return total / float64(n)
}
func main() {
	var balita arrbalita
	var n int
	fmt.Print("Banyak balita: ")
	fmt.Scan(&n)

	hitung(&balita, n)
	rata_rata:= rata(&balita,n)
	fmt.Printf("%.2f\n", rata_rata)
}