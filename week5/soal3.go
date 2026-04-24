//Fathi Satriani Al_109082530029
package main

import "fmt"

func pangkat(x, y int) int{
	if y == 0{
		return 1
	}
	hasil:= x * pangkat(x, y-1)
	return hasil
}
func main() {
	var x, y int

	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}