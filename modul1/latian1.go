package main

import "fmt"
func loss (z int) int{
	theloss := (z + z)
	return theloss
}

func main() {
	persen  := 0
	for i:= 0; i < 9; i++{
		if i == 0 || i%3 == 0{
			fmt.Println(i, "BINGO")
		}else if i < 3{
			fmt.Println(i, "train active", persen, "% accuracy, but", loss(i+1))
		}else{
			fmt.Println(i, "trains active", persen, "% accuracy, but", loss(i))
		}
		persen += 6
	}
}

