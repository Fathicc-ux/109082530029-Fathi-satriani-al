# <h1 align="center">Laporan Praktikum Modul 10 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Cari Min Max ]
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.

#### Soal1.go

```go
package main

import "fmt"

func minmax(arr [1000] float64, n int){
	for i:=0; i<n; i++{
		fmt.Print("berat Kelinci ", i+1, ": ")
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	max := arr[0]
	for i:=1; i<n; i++{
		if arr[i] < min{
			min = arr[i]
		}else if arr[i] > max{
			max = arr[i]
		}
	}
	fmt.Println("berat kelinci terkecil: ", min)
	fmt.Println("berat kelinci terbesar: ", max)
}
func main() {
	var arr [1000] float64
	var n int
	
	fmt.Print("Banyak kelinci: ")
	fmt.Scan(&n)

	minmax(arr, n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas10/output_tugas10/Soal1.png)
[penjelasan]
Program digunakan untuk mencari nilai min max dari sebuah input, input pertama menyatakan banyaknya kelinci input kedua adalah berat tiap kelinci. (berat tiap kelinci dalam index) setelah itu buat perulangan yang didalamnya ada percabangan untuk menentukan nilai min max, untuk perulangan saya mulai dari 1, dan percabangannya saya bandingkan dengan index [0] yang dimana jika nilai index [i] < nilai index [0] maka nilai index [0] akan digantikan dengan nilai index [i]

### 2. [Mencari jumlah total dan rata rata tiap wadah ikan]
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah. 
#### Soal2.go

```go
package main

import "fmt"

func main(){
	var ikan [1000] float64
	var x, y int
	var total, jumlah float64
	fmt.Print("x, y: ")
	fmt.Scan(&x, &y)

	for i:=0; i<x; i++{
		fmt.Scan(&ikan[i])
	}
	for i:=0; i<x; i+=y{
		total = 0
		jumlah = 0
		for j:= i; j<i+y && j<x; j++{
			total = total + (ikan[j])
			jumlah++
		}
		fmt.Print(total, " ")
	}
fmt.Println()
	for i:=0; i<x; i+=y{
		total = 0
		jumlah = 0
		for j:= i; j<i+y && j<x; j++{
			total = total + (ikan[j])
			jumlah++
		}
		fmt.Print(total/jumlah, " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas10/output_tugas10/Soal2.png)
[penjelasan]
Ptogram digunakan untuk mencari total jumlah ikan tiap wadah, yang dimana isi ikan dalam wadah beratnya harus urut dengan apa yang diinput. Awal mula program meminta input baris pertama x,y "x" sebagai banyaknya ikan dan "y" jumlah ikan yang akan dimasukan dalam wadah, baris kedua input berat tiap ikan sebanyak x, lalu masuk ke perulangan dimulai dari 0 dan perulangan akan menambah sebanyak  nilai y untuk iterasi berikutnya sebagai jumlah wadah. Jumlah wadah ditentukan dari nilai x, y diawal keluran 2 baris, baris pertama menyatakan jumlah ikan tiap wadah dan baris kedua rata rata berat dalam tiap wadahnya

### 3. [Mencari nilai terberat dan teringan balita]
Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya. 
#### Soal3.go

```go
package main

import "fmt"

type arrbalita [100] float64

func hitung(arrberat *arrbalita, n int){
	for i:=0; i<n; i++{
		fmt.Print("Berat balita ke-", i+1," : ")
		fmt.Scan(&arrberat[i])
	}
	min := arrberat[0]
	max := arrberat[0]
	for i:=1; i<n; i++{
		if arrberat[i] < min{
			min = arrberat[i]
		}else if arrberat[i] > max{
			max = arrberat[i]
		}
	}
	fmt.Println("berat balita terkecil: ", min, "kg")
	fmt.Println("berat balita terbesar: ", max, "kg")
}
func rata(arrberat *arrbalita, n int)float64{
	total:= 0.0
	for i:=0 ;i<n; i++{
		total = total + arrberat[i]
	}
	return total / float64(n)
}
func main() {
	var balita arrbalita
	var n int
	fmt.Print("Banyak balita: ")
	fmt.Scan(&n)

	hitung(&balita, n)
	rata_rata:= rata(&balita,n)
	fmt.Printf("%.2f\n", rata_rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas10/output_tugas10/Soal3.png)
[penjelasan]
Program ini digunakan untuk mencari nilai min max, programnya hampir sama seperti nomor 1 tapi ini ada tambahan untuk hitung rata rata, yang dimana berat balita yang kita input ditambahkan lalu hasilnya dibagi banyaknya balita

