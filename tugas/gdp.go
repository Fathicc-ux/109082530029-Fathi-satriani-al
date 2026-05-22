package main

import "fmt"

func kpita(gdp, penduduk float64) float64{
	return gdp/penduduk
}
func rata(n int){
	var total float64 = 0
	for i:=1; i<=n; i++{
		var gdp, penduduk  float64
		fmt.Print("GDP negara ", i, ": ")
		fmt.Scan(&gdp)

		fmt.Print("penduduk negara ", i, ": ")
		fmt.Scan(&penduduk)
		hasil:= kpita(gdp,penduduk)
		fmt.Println("GDP perkapita: ", hasil)
		total = total + hasil
	 }
	 rata:= total / float64(n)
	 fmt.Println("Rata rata GDP perkapita: ", rata)
}
func main() {
	var n int
	fmt.Print("Banyak negara: ")
	fmt.Scan(&n)
	 
	rata(n)
}