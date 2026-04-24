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
Program ini untuk mengetahui posisi sebuah titik lingkaran apakah ada titi di dalam lingkaran 1 dan 2, titik di dalam lingkaran 1, titik di dalam lingkaran 2 atau titik diluar lingkaran 1 dan 2. Program ini menggunakan tipe data bentukan yang dimana di dalamnya menyimpan beberapa variabel dengan tipe data yang berbeda, program meminta input titik koordinat dan radius lingkaran 1 dan 2, dan titik sembarang. Lalu masuk ke fungsi jarak untuk menghitung jarak titik (a, b) dan (c, d) dengan rumus: 
√((a − c)*(a − c)) + ((b − d)*(b − d)), lalu masuk ke fungsi dalam untuk mengecek apakah titik di dalam lingkaran 1 dan 2 atau diluar lingkaran 1 dan 2. Setelah itu masuk e fungsi main di if untuk menentukan apakah titik berada di dalam lingkaran 1 dan 2 atau diluar lingkaran 1 dan 2.

### 2. [Array]
Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat
menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.
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
Sebuah program array digunakan untuk menampilkan beberapa infomasi seperti menampilkan seluruh isi array, index ganjil genap, kelipatan dengan input x, hapus index lalu tampilkan setelah dihapus, rata rata, standar deviasi, frekuensi. Disini saya menggunakan beberapa func prosedur untuk membuat program ini, setiap prosedur digunakan untuk menampilkan informasi diatas.


### 3. [Mencari pemenang pertandingan]
Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian
program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan
dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan. 
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
Program digunakan untuk mencari pemenang pertandingan sepak bola, array dalam program ini digyunakan untuk menyimapn daftar pemenang lalu ditampilkan hasilnya. Program meminta input nama klub, dan jumlah gol dari setiap pertandingan, input akan berhenti ketika gol dari kedua tim ada yang kurang dari 0 (-). Program ini menggunakan percabangan untuk menentukan klub pemenang dari setiap pertandingan, lalu hasil klub pemenang akan di simpan dalam sebuah array dan ditampilkan di hasil akhir dari setiap pertandingan.


### 4. [Palindrom]
Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom. Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi
untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama. 
#### Soal4.go

```go
package main

import "fmt"

func array(NMAX *[127] rune, n *int){
	var kata string
	i:=0
	fmt.Print("isi: ")
	for {
		fmt.Scan(&kata)
		if kata == "."{
			break
		}
		(*NMAX) [i] =  rune(kata[0])
		i++
	}
	*n = i
}
func balikan(NMAX *[127] rune, n int){
	fmt.Print("balik: ")
	for i:=n-1; i>=0; i--{
		 fmt.Print(string((*NMAX)[i]), " ")
	}
}
func palindrom(NMAX *[127] rune, n int) bool{
	for i:=0; i<n/2; i++{
		if(*NMAX)[i] != (*NMAX)[n-1-i]{
			return false
		}
	}
	return  true
}	
func main() {
	var NMAX [127] rune
	var n int

	array(&NMAX, &n)
	fmt.Print("cetak array: ")
for i := 0; i < n; i++ {
	fmt.Print(string(NMAX[i]), " ")
}
fmt.Println()
	 balikan(&NMAX, n)
fmt.Println()
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
Sebuah array digunakan untuk menampung sekumpulan karakter, Program diminta input sebanyak banyak nya karakter dan akan berhenti saat input '.' kemudian program mencetak input tersebut. Fungsi array digunakan untuk input, lalu mencetaknya di fungsi main setelah itu fungsi balikan digunakan untuk membalikan hasil cetak. Lalu gunakan fungsi palindrom apakah karakter yag diceak dengan yang dibalik hasilnya tetap sama, jika sama maka dianggap palindrom dan kalau beda bukan.