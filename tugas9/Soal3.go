package main

import "fmt"
func hitung(klubA, klubB string){
	var goalA, goalB int
	var n int = 0
	var menang [100] string
	jumlah := 1 
	
	for{
	fmt.Print("Skor pertandingan ", jumlah, ": " )
	fmt.Scan(&goalA, &goalB)
	if goalA < 0 || goalB < 0{
		break
	}

	if goalA > goalB{
		menang[n] = klubA
		n++
	}else if goalA < goalB{
		menang[n] = klubB
		n++
	}
	jumlah++
}
	for i:= 0; i<n; i++{
		fmt.Println("Hasil ", i+1, ": ", menang[i])
	}
}
func main(){
	 var klubA, klubB string
	
	fmt.Print("Klub A: ")
	 fmt.Scan(&klubA)
	 fmt.Print("Klub B: ")
	 fmt.Scan(&klubB)
	
	 hitung(klubA, klubB)
}