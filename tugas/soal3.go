package main

import "fmt"

func main() {
	var parsel, biaya, Rp, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)

	kg := parsel / 1000
	detail_berat:= parsel % 1000
	fmt.Println("Detail berat: ", kg, "kg", "+" ,detail_berat, "gram")
	Rp = kg * 10000

	if parsel > 10000{
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