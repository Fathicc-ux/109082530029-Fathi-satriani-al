# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Fathi satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Faktorial, Permutasi, Kombinasi]
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p) 
Masukan: terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d. 
Keluaran: terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### Soal1.go

```go
package main

import "fmt"

func faktorial(n int)int{
	hasil:= 1
	for i:=1; i<=n; i++{
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int)int{
	hasil := faktorial(n)/faktorial(n-r)
	return hasil
}
func kombinasi(n, r int)int{
	hasil := faktorial(n) / ((faktorial(r) * faktorial(n-r)))
	return hasil
}
func main () {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println("Baris 1: ", permutasi(a, c), kombinasi(a, c))
	fmt.Println("Baris 2: ", permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas3/output_tugas3/Soal1.png)
[penjelasan]
blablablablablabla


### 2. [Fungsi Komposisi]
Diberikan tiga buah fungsi matematika yaitu f (x) = x * x, g (x) = x − 2 dan h (x) = x +1 . Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function. 
Masukan: terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. 
Keluaran: terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!
#### Soal2.go

```go
package main

import "fmt"

func fx(a int)int{
	hasil:= a * a
	return hasil
}
func gx(a int)int{
	hasil:= a -2
	return hasil
}
func hx(a int)int{
	hasil:= a + 1
	return hasil
}
func main() {
	var a, b, c int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas3/output_tugas3/Soal2.png)
[penjelasan]
program 


### 3. [Titik Lingkaran]
Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. 
Masukan: terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat. 
Keluaran: berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### Soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c ,d float64)float64{
	hasil:= math.Sqrt((a-c)*(a-c) + (b - d)*(b-d))
	return hasil
}
func dalam(cx, cy, r, x, y float64)bool{
	if jarak(x, y, cx, cy) <= r{
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

	daling1:= dalam(float64(a1), float64(b1), float64(r1), float64(x), float64(y))
	daling2:= dalam(float64(a2), float64(b2), float64(r2), float64(x), float64(y))

	if daling1 && daling2{
		fmt.Println("Titik di dalam lingkaran 1 & 2")
	}else if daling1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if daling2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 & 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas3/output_tugas3/Soal3.png)
[penjelasan]
program 