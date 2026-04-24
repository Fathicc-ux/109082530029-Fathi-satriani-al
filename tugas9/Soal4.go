package main

import "fmt"


// func array(arr *[256] int, n int){
// 	for i:=0; i<n; i++{
// 		fmt.Scan(&arr[i])
// 	}
// }
// func revrese(arr *[256] int, n int){
// 	var a [256] int
// 	for i:=0; i<n; i++{
// 		a[i] = arr[n-1-i]
// 	}
// 	for i:= 0; i<n; i++{
// 		arr[i] = a[i] 
// 	}
// }
// func palindrom(arr [256] int, n int) bool{
// 	var a [256] int
// 	for i:=0; i<n; i++{
// 		a[i] = arr[n-1-i]
// 	}
// 	for i:= 0; i<n; i++{
// 		if arr[i] != a[i]{
// 			return false
// 		}
// 	}
// 	return  true
// }	
// func main() {
// 	var a [256] int
// 	var n int
// 	fmt.Print("n: ")
// 	fmt.Scan(&n)

// 	array(&a, n)
// 	fmt.Println("isi array:")
// for i := 0; i < n; i++ {
// 	fmt.Print(a[i], " ")
// }
// fmt.Println()
// 	if palindrom(a, n){
// 		fmt.Println("Palindrom")
// 	}else{
// 		fmt.Print("bukan")
// 	}
// 	revrese(&a, n)
// }