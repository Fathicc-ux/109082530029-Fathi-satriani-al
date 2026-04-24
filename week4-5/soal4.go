//Fathi Satriani Al_109082530029

package main

import "fmt"

func zoomin(w, h,skala  int){
		w = w*skala
		h = h*skala
		fmt.Print(w, h)
}
func zoomout(w, h, skala int){
	w = w/skala
	h = h/skala
	fmt.Print(w, h)
}
func main() {
	var w, h, skala int
	var inout string
	fmt.Scan(&w, &h)
	fmt.Scan(&inout, &skala)

	if inout == "+"{
		zoomin(w, h, skala)
	}else{
		zoomout(w, h, skala)
	}
}