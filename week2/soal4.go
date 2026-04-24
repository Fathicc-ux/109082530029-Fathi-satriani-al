//Fathi Satriani Al_109082530029
package main

import "fmt"

func fibo(n int) int{
	if n == 1{
		return 0
	}else if n == 2 {
		return 1
}
		return fibo(n-1) + fibo(n-2)
	}

func main() {
	var n int
	fmt.Print("input: ")
	fmt.Scan(&n)
	hasil:= 0
	for i:= 1; i<=n; i++{
    hasil = hasil + fibo(i)
	}
	fmt.Print("output: ", hasil)
}