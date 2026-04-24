//Fathi Satriani Al_109082530029
package main
import "fmt"

func bintang(n int){
    if n == 0{
        return
    }
    bintang(n-1)
    for i:=1; i<=n; i++{
        fmt.Print("*")
    }
    fmt.Println( )
}
func main() {
    var n int
    fmt.Print("Angka: ")
    fmt.Scan(&n)
    
    bintang(n)
}