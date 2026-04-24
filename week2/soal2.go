//Fathi Satriani Al_109082530029
package main

import "fmt"

func Segitiga(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func main() {
	var a, b, c int
	fmt.Print("Masukkan sisi: ")
	fmt.Scan(&a, &b, &c)

	if Segitiga(a, b, c) {
		fmt.Println("segitiga")
	} else {
		fmt.Println("bukan segitiga")
	}
}