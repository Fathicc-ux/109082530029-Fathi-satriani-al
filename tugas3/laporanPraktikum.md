# <h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

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
Program di gunakan untuk menghitung permutasi dan kombinasi, mahaisswa harus menyelesaikan subprogram dari soal. Ada 3 Function Faktorial, Permutasi dan kombinasi Function faktorial di gunakan untuk menghitung hasil dari permutasi dan kombinasi, di mana rumus permutasi dan kombinasi membutuhkan bilangan faktorial. 
Program input 4 bilangan a, b, c, d, lalu masuk ke function faktorial contoh input: 5 10 3 10 menjadi 5! 10! 3! 10! program ingin mencetak function permutasi dan kombinasi pada baris 1 dan 2. Masuk ke Permutasi (baris 1 meminta a dengan c, baris 2 b dengan d) pada rumus permutasi di program ibaratkan n = a dan r = c (begitupun b dengan d). Permutasi 5!/(5-3)! = 5!/2! = 120/2 = 60 (Permutasi). 
Lanjut kombinasi dengan rumus n!/(r! * (n-r)!) = 5!/(3! * (5-3)!) = 5!/ (3! * 2!) = 120 / 12 = 10

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
program untuk menghitung fungsi f(x) g(x) h(x), Input 3 bilangan yang nantinya akan masuk di ketiga function tadi.Input 7 2 10 ibaratkan sebagai (a b c) f(g(h(a))), g(h(f(b))), h(f(g(c)))
7 masuk function h(x) = 7+1 = 8, g(x) = 8-2 = 6, h(x)=6*6= 36.
2 masuk function f(x) = 2*2 = 4, h(X) = 4+1 = 5, g(x) = 5-2 = 3.
10 masuk function g(x) = 10-2 = 8, f(X) = 8*8 = 64, h(x) = 64+1 = 65.
Input 7 2 10 Outputnya 36 3 65


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
program di gunakan untuk mengecek tiitk 2 lingkaran apakah di dalam atau diluar. Terdapat 2 function Jarak dan Dalam function jarak digunakan untuk menghitung titik ousat dengan x,y sedangkan dalam untuk memberikan nilai kebenaran apakah titik lingkaran diluar atau di dalam. Masuk function Jarak Contoh input Baris 1 = 1 2 3 (1, 2 titik pusat lingkran 1, 3 = radius), Baris 2 = 4 5 6 (4, 5 titik pusat lingkaran 2, 6 = radius), Baris 3 = 7 8 (7, 8 titik yang diuji dengan pusat lingkaran 1 & 2).
Masuk function jarak hitung lingkaran 1 = akar(7-1 * 7-1) + (8-2 * 8-2) = akar(36+36) = akar72 = 
8.4, karena jarak lebih dari radius berarti titik lingkaran 1 ada di luar (8.4 > 3)
Hitung lingkaran 2 = akar(7-4 * 7-4) + (8-5 * 8-5) = akar(9+9) = akar18 = 4.2. Karena jarak kurang dari radius berarti titik lingkaran 2 ada di dalam (4.2 < 6). Masuk function dalam, hasil hitung lingkaran 1 menunjukan titik diluar dan lingkaran 2 di dalam, karena yang benar cuman 1 yaitu lingkaran 2 maka outputnya : "Titik di dalam lingkaran 2"