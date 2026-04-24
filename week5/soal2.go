//Fathi Satriani Al_109082530029
package main

import "fmt"

func faktor(i, n int){
	if i > n{
		return
	}
	for i= 1; i<=n; i++{
		if n%i == 0{
			fmt.Print(i, " ")
		}
	}
	faktor(i+1, n)
}
func main() {
	var n int

	fmt.Scan(&n)
	
	faktor(1, n)
}