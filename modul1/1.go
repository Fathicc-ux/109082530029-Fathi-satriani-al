package main
import "fmt"
type arrayKata [8]string
func isiArray(arrKata *arrayKata) {
	/* I.S. Terdefinisi array untuk menampung string huruf
	   F.S. arrKata berisi 8 huruf unik */

	for i := 0; i < 8; i++ {
		fmt.Scan(&arrKata[i])
	}
}

func arrayToString(arrKata arrayKata) string {
	/* mengembalikan array hasil konversi 8 huruf menjadi sebuah string */

	var hasil string

	for i := 0; i < 8; i++ {
		hasil = hasil + arrKata[i]
	}

	return hasil
}

func main() {
    
	var arrKata arrayKata
	// var hasil string

	isiArray(&arrKata)

	fmt.Println(arrKata, "diubah menjadi :", (arrayToString(arrKata)))
}