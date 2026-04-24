//Fathi Satriani Al_109082530029
package main

import "fmt"
func koversi(n float64){
	reamur:= (4.0 / 5.0) * n
	fahrenhait:= (n * (9.0 / 5.0) + 32)
	kelvin:= n + 273.15

	fmt.Print("reamur: ", reamur, " fahrenhait: ", fahrenhait, " kelvin: ", kelvin)
}
func main() {
	var cel float64
	fmt.Print("satuan celcius: ")
	fmt.Scan(&cel)

	koversi(cel)
}