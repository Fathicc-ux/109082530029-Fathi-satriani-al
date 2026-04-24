# <h1 align="center">Laporan Praktikum Modul 9 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Cek lingkaran dengan tipe bentukan ]
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan
koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### Soal1.go

```go
package main

import (
	"fmt"
	"math"
)
type titik struct {
	x, y float64
}

type lingkaran struct {
	x, y, r float64
}
func jarak(p, q titik) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}
func dalam(c lingkaran, p titik) bool {
	if jarak(titik{c.x, c.y}, p) <= c.r {
		return true
	}
	return false
}
func main () {
	var a1, b1, r1 int
	var a2, b2, r2 int
	var x, y int

	fmt.Scanln(&a1, &b1, &r1)
	fmt.Scanln(&a2, &b2, &r2)
	fmt.Scanln(&x, &y)

	daling1 := dalam(
		lingkaran{float64(a1), float64(b1), float64(r1)},
		titik{float64(x), float64(y)},
	)

	daling2 := dalam(
		lingkaran{float64(a2), float64(b2), float64(r2)},
		titik{float64(x), float64(y)},
	)

	if daling1 && daling2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if daling1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if daling2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas9/output_tugas9/Soal1.png)
[penjelasan]
Program 

### 2. [Pola bintang]
Buatlah 
#### Soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func seluruh(arr [256] int, n int) [256] int{
	
	 for i:= 0; i<n; i++{
		fmt.Scan(&arr[i])
	 }
	 fmt.Print("nilai seluruh: ")
	 for i:= 0; i<n; i++{
		fmt.Print(arr[i], " ")
	 }
	 fmt.Println()
	 return arr
}
	
func ganjil(arr [256] int, n int){
	fmt.Print("ganjil: ")
	for i:= 0; i<n; i++{
		if i%2 == 1{
		fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func genap(arr [256] int, n int){
	fmt.Print("genap: ")
	for i:= 0; i<n; i++{
		if i%2 == 0{
		fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func kelipatan(arr [256]int, n int) {
	var x int
	fmt.Print("x: ")
	fmt.Scan(&x)

	fmt.Print("kelipatan: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}
func hapus(arr [256]int, n int) ([256] int, int){
	var index int
	fmt.Print("hapus index: ")
	fmt.Scan(&index)

	for i:= index; i<n-1; i++{
		arr[i] = arr[i+1]
	}
	n --
	fmt.Print("hasil hapus: ")
	for i:= 0; i<n; i++{
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	return arr, n
}
func rata(arr [256] int, n int){
	jumlah:= 0
	for i:= 0; i < n; i++{
		jumlah+= arr[i]
	}
	rata:= jumlah / n
	fmt.Println("rata rata: ", rata)
}
func simpangan_baku(arr [256] int, n int){
	var total float64
	var jumlah int
	for i:= 0; i < n; i++{
		jumlah+= arr[i]
	}
	rata:= float64(jumlah) / float64(n)

	for i:=0; i<n; i++{
		selisih:= float64(arr[i]) - (rata)
		total+= (selisih) * (selisih)
	}
	fmt.Print("simpangan baku: ", math.Sqrt(float64(total)/ float64(n)))
	fmt.Println()
}
func frekuensi(arr [256] int, n int){
	var x, hitung int
	fmt.Print("frek: ")
	fmt.Scan(&x)
	for i:=0 ;i<n; i++{
		if arr[i] == x{
			hitung++
		}
	}
	fmt.Print("frekuensi: ", hitung)
}
func main () {
	var n int
	var arr [256] int
	fmt.Print("banyak n: ")
	fmt.Scan(&n)
	
	arr = seluruh(arr, n)
	ganjil(arr, n)
	genap(arr, n)
	kelipatan(arr, n)
	arr, n = hapus(arr, n)
	rata(arr, n)
	simpangan_baku(arr, n)
	frekuensi(arr, n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas9/output_tugas9/Soal2.png)
[penjelasan]
Program 


### 3. [Mencari faktor dengan rekrusif]
Buatlah 
#### Soal3.go

```go
package main

import "fmt"
func hitung(klubA, klubB string){
	var goalA, goalB int
	var n int = 0
	var menang [100] string
	jumlah := 1 
	
	for{
	fmt.Print("Skor pertandingan ", jumlah, ": " )
	fmt.Scan(&goalA, &goalB)
	if goalA < 0 || goalB < 0{
		break
	}

	if goalA > goalB{
		menang[n] = klubA
		n++
	}else if goalA < goalB{
		menang[n] = klubB
		n++
	}
	jumlah++
}
	for i:= 0; i<n; i++{
		fmt.Println("Hasil ", i+1, ": ", menang[i])
	}
}
func main(){
	 var klubA, klubB string
	
	fmt.Print("Klub A: ")
	 fmt.Scan(&klubA)
	 fmt.Print("Klub B: ")
	 fmt.Scan(&klubB)
	
	 hitung(klubA, klubB)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas9/output_tugas9/Soal3.png)
[penjelasan]
Program


### 4. [Mencari faktor dengan rekrusif]
Buatlah 
#### Soal4.go

```go
package main

import "fmt"

func array(NMAX *[127] rune, n int){
	for i:=0; i<n; i++{
		var kata string
		fmt.Scan(&kata)
		NMAX [i] =  rune(kata[0])
	}
}
func revrese(NMAX *[127] rune, n int){
	var a [127] rune
	for i:=0; i<n; i++{
		a[i] = NMAX[n-1-i]
	}
	for i:= 0; i<n; i++{
		NMAX[i] = a[i] 
	}
}
func palindrom(NMAX *[127] rune, n int) bool{
	var a [127] rune
	for i:=0; i<n; i++{
		a[i] = (NMAX[n-1-i])
	}
	for i:= 0; i<n; i++{
		if NMAX[i] != a[i]{
			return false
		}
	}
	return  true
}	
func main() {
	var NMAX [127] rune
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)

	array(&NMAX, n)
	fmt.Println("ISI ARRAY:")
for i := 0; i < n; i++ {
	fmt.Print(string(NMAX[i]), " ")
}
fmt.Println()
revrese(&NMAX, n)
	if palindrom(&NMAX, n){
		fmt.Println("Palindrom")
	}else{
		fmt.Print("bukan")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas9/output_tugas9/Soal4.png)
[penjelasan]
Program 