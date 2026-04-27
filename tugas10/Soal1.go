package main

import "fmt"

func minmax(arr [1000] float64, n int){
	for i:=0; i<n; i++{
		fmt.Print("berat Kelinci ", i+1, ": ")
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	max := arr[0]
	for i:=1; i<n; i++{
		if arr[i] < min{
			min = arr[i]
		}else if arr[i] > max{
			max = arr[i]
		}
	}
	fmt.Println("berat kelinci terkecil: ", min)
	fmt.Println("berat kelinci terbesar: ", max)
}
func main() {
	var arr [1000] float64
	var n int
	
	fmt.Print("Banyak kelinci: ")
	fmt.Scan(&n)

	minmax(arr, n)
}