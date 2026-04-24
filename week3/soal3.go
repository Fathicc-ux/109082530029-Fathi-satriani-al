//Fathi Satriani Al_109082530029
package main

import "fmt"
func pengurutan(x, y int){
	if x > y{
		fmt.Print(y, x)
	}else if x < y{
		fmt.Print(x, y)
	}else{
		fmt.Print()
	}
}
func main() {
	var x, y int

	fmt.Scan(&x, &y)
	 
	pengurutan(x, y)
}