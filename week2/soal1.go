//Fathi Satriani Al_109082530029
package main

import "fmt"

func hitungjam(jam, menit, detik int)int{
	return jam * 3600 + menit * 60 + detik
}
func main() {
	var jam, menit, detik int

	fmt.Scan(&jam, &menit, &detik)

	fmt.Print(hitungjam(jam, menit, detik), " detik")
}