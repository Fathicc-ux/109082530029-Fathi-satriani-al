package main

import "fmt"

func main(){
	var ikan [1000] float64
	var x, y int
	var total, jumlah float64
	fmt.Print("x, y: ")
	fmt.Scan(&x, &y)

	for i:=0; i<x; i++{
		fmt.Scan(&ikan[i])
	}
	for i:=0; i<x; i+=y{
		total = 0
		jumlah = 0
		for j:= i; j<i+y && j<x; j++{
			total = total + (ikan[j])
			jumlah++
		}
		fmt.Print(total, " ")
	}
fmt.Println()
	for i:=0; i<x; i+=y{
		total = 0
		jumlah = 0
		for j:= i; j<i+y && j<x; j++{
			total = total + (ikan[j])
			jumlah++
		}
		fmt.Print(total/jumlah, " ")
	}
}