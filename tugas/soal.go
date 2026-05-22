package main

import "fmt"

func parsel(berat int) (kg, detail_berat, biaya, total int){
	
	kg  = berat / 1000
	detail_berat = berat % 1000
	
	Rp := kg * 10000

	if berat > 10000{
		biaya = 0
	}else  if detail_berat >= 500{
		biaya = 5 * detail_berat 
	}else {
		biaya = 15 * detail_berat
	}
	total = Rp + biaya
	return kg, detail_berat, biaya, total
}
func main() {
	var berat int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg, detail_berat, biaya, total:= parsel(berat)

	fmt.Println("Detail berat: ", kg, "kg", "+" ,detail_berat, "gram")
	fmt.Println("Detail biaya: ", "Rp. ", kg*10000 ," + ", "Rp. ", biaya)
	fmt.Println("total biaya Rp.", total)
}
