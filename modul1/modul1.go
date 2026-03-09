package main
import "fmt"

func area (z int) int{
	thearea := z * z
	return thearea
}
func main() {
	height:=8
	vol:= height * area(1)
	fmt.Print(vol)
}