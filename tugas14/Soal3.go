package main

import "fmt"

func median(arr [1000]int, n, data int){
	selesai:= false
	fmt.Print("Masukan Angka: ")
	for !selesai{
		fmt.Scan(&n)

		if n == -5313{
			selesai = true
		}else if n != 0{
			arr[data] = n
			data++
		}else if n  == 0{
			for i:=0; i<data-1; i++{
			min := i
			for j:=i+1; j<data; j++{
			if arr[min] > arr[j]{
				min = j
						}
					}
			temp:= arr[min]
			arr[min] = arr[i]
			arr[i] = temp
				}
				if data % 2 == 1{
					median:= data /2
					fmt.Println("Median: ", arr[median])
				}else{
					median:= (arr[data/2-1] + arr[data] / 2) / 2
					fmt.Println("Median: ", arr[median])
				}
			}
	}
}
func main() {
	var arr [1000]int
	var n, data int
	
	median(arr, n, data)
}