package main

import "fmt"
func selection(m [1000] int, rumah int, daerah int){
	fmt.Print("Jumlah daerah: ")
	fmt.Scan(&daerah)
	for i:=0; i<daerah; i++{
		fmt.Print("Jumlah rumah daerah ke-", i+1,": ")
		fmt.Scan(&rumah)

	fmt.Print("Nomor rumah: ")
	for i:=0; i<rumah; i++{
		fmt.Scan(&m[i])
	}

	for i:=0; i<rumah-1; i++{
		min := i
		for j:=i+1; j<rumah; j++{
			if m[min] > m[j]{
				min = j
			}
		}
		temp:= m[min]
		m[min] = m[i]
		m[i] = temp
		}
		fmt.Print("Urutan Nomor rumah daerah ke-",i+1 ," yang dikunjungi: ")
		for i:=0; i<rumah; i++{
			fmt.Print(m[i], " ")
		}
		fmt.Println()
	}
}
func main() {
	var m [1000] int
	var daerah int
	var rumah int

	selection(m, daerah, rumah)
}