//Fathi Satriani Al_109082530029
package main

import "fmt"
func hitungdiskon(bel int, mem bool) float64{
	if bel > 100000 && mem == true{
		return  float64(bel) * 0.10
	}else if bel > 100000 && mem == false {
		return  float64(bel) * 0.05
	}else{
		return 0
	}
}

func main() {
	var belanja int
	var member bool

	fmt.Scan(&belanja, &member)
	diskon:= hitungdiskon(belanja, member)
	fmt.Print(belanja - int(diskon))
}