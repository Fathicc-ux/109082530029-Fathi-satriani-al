package main

import "fmt"

func parsel(berat int){
	var biaya, Rp, total, detail_berat int

	kg := berat / 1000
	detail_berat = berat % 1000
	fmt.Println("Detail berat: ", kg, "kg", "+" ,detail_berat, "gram")
	Rp = kg * 10000

	if berat > 10000{
		biaya = 0
	}else  if detail_berat >= 500{
		biaya = 5 * detail_berat 
	}else {
		biaya = 15 * detail_berat
	}
	fmt.Println("Detail biaya: ", "Rp. ", Rp, " + ", "Rp. ", biaya)

	total = Rp + biaya

	fmt.Print("Total biaya: ","Rp. ", total)
}
func main() {
	var berat int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	parsel(berat)
}