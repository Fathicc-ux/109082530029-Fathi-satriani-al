//Fathi Satriani Al_109082530029
package main

import "fmt"

func volume(r, t float64) float64{
	hasil:= 3.14 * (r * r) * t
	return hasil
}
func massa(r, t, p float64) float64{
	hasil:= volume(r, t) * p
	return hasil
}
func display(mkiri, mkanan float64){
	if mkiri == mkanan{
		fmt.Print("Balance")
	}else if mkiri > mkanan{
		fmt.Printf("%.3f", mkiri - mkanan)
	}else{
		fmt.Printf("%.3f", mkanan - mkiri)
	}
}
func main(){
	var r, tkiri, pkiri, tkanan, pkanan float64
	fmt.Scan(&r)
	fmt.Scan(&tkiri, &pkiri)
	fmt.Scan(&tkanan, &pkanan)

	(volume(r, tkiri))
	(volume(r, tkanan))
	(massa(r, tkiri, pkiri))
	(massa(r, tkanan, pkanan))

	mkiri:= massa(r, tkiri, pkiri)
	mkanan:= massa(r, tkanan, pkanan)

	display(mkiri, mkanan)
}