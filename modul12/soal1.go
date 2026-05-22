package main

import "fmt"


func main() {
	n:=9
	tim := [9]string {"RRQ", "EVOS", "BTR", "ONIC", "NAVI", "DEWA", "AE", "GEEK", "TILD"}
	poin := [9]int {1, 8, 7, 12, 6, 8, 7, 7, 8}

	for i:= 0; i<9; i++{
		fmt.Println("TIM: ", tim[i], " dengan point: ", poin[i])
	}
	
}