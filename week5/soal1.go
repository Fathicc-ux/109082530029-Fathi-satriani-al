//Fathi Satriani Al_109082530029
package main

import "fmt"
func ganjil(i, n int){
	if i > n{
		return
	}
	for i = 1; i<=n; i++{
		if i%2 == 1{
			fmt.Print(i, " ")
		}
	}
	ganjil(i+1, n)
}
func main() {
	var n int

	fmt.Scan(&n)
	
	ganjil(1, n)
}