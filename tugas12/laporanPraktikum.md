# <h1 align="center">Laporan Praktikum Modul 12 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [rumahkerabat]
Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya
Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program
rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut
membesar menggunakan algoritma selection sort.

Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat
Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m <
1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan
rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-
masing daerah. 

#### Soal1.go

```go
package main

import "fmt"
func selection(m [1000] int, rumah int, daerah int){
	fmt.Print("Jumlah daerah: ")
	fmt.Scan(&daerah)
	for i:=0; i<daerah; i++{
		fmt.Print("Jumlah rumah daerah ke-", i+1,": ")
		fmt.Scan(&rumah)

	fmt.Print("Nomor rumah: ")
	for i:=0; i<rumah; i++{
		fmt.Scan(&m[i])
	}

	for i:=0; i<rumah-1; i++{
		min := i
		for j:=i+1; j<rumah; j++{
			if m[min] > m[j]{
				min = j
			}
		}
		temp:= m[min]
		m[min] = m[i]
		m[i] = temp
		}
		fmt.Print("Urutan Nomor rumah daerah ke-",i+1 ," yang dikunjungi: ")
		for i:=0; i<rumah; i++{
			fmt.Print(m[i], " ")
		}
		fmt.Println()
	}
}
func main() {
	var m [1000] int
	var daerah int
	var rumah int

	selection(m, daerah, rumah)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas12/output_tugas12/Soal1.png)
[penjelasan]
Program selection sort ini digunakan untuk mengurutkan nilai dari yang terkecil ke terbesar hampir sama seperti mencari nilai ekstrem namun bedanya selection sort bukan sekedar mencari nilai dari terkecil ke terbesar tapi menukar posisi data juga. Pada program ini ada 3 input, pertama input banyak daerah yang dikunjungi, kedua jumlah rumah pada daerah tersebut, 3 nomor rumah pada daerah tersebut. Awal input jumlah daerah setelah itu input banyaknya rumah paada daerah itu lalu input juga nomor rumah sebanyak jumlah rumah, setelah itu eksekusi sorting menggunakan for loop dimana membandingan kedua index array, pada perulangan pertama bandingkan idx:min dan idx:j kalau TRUE maka nilai min = j setelah itu melakukan penukaran posisi data

### 2. [kerabat dekat]
Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu
diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena
nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.
Format Masukan masih persis sama seperti sebelumnya.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah. 
#### Soal2.go

```go
package main

import "fmt"
func Sorting(m [1000] int, rumah int, daerah int){
	fmt.Print("Jumlah daerah: ")
	fmt.Scan(&daerah)
	for i:=0; i<daerah; i++{
		fmt.Print("Jumlah rumah daerah ke-", i+1,": ")
		fmt.Scan(&rumah)

	fmt.Print("Nomor rumah: ")
	for i:=0; i<rumah; i++{
		fmt.Scan(&m[i])
	}

	for i:=0; i<rumah-1; i++{
		min := i
		for j:=i+1; j<rumah; j++{
			if m[min] > m[j]{
				min = j
			}
		}
		temp:= m[min]
		m[min] = m[i]
		m[i] = temp
		}
		fmt.Print("Urutan Nomor rumah daerah ke-",i+1 ," yang dikunjungi: ")
		for i:=0; i<rumah; i++{
			if m[i] % 2 == 1{
				fmt.Print(m[i], " ")
			}
		}
		for i:=rumah-1; i>=0; i--{
			if m[i] % 2 == 0{
				fmt.Print(m[i], " ")
			}
		}
		fmt.Println()
	}
}
func main() {
	var m [1000] int
	var daerah int
	var rumah int

	Sorting(m, daerah, rumah)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas12/output_tugas12/Soal2.png)
[penjelasan]
Program tersebut hampir sama dengan soal nomor 1 bedanya kalau nilai idx ganjil maka akan diurutkan terkecil ke terbesar, jika nilai idx genap maka terbesar ke terkecil, input daerah 1 campuran ganjil genap, daerah 2 ganjil, daerah 3 genap. Bedanya dengan kode program nomor 1 ada dibagian cetaknya saja

### 3. [Mencari nilai median]
Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan
tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan
sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu
problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba
untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?
"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data
genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua
data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke
bawah."
Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah
terbaca, jika data yang dibaca saat itu adalah 0.
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
Keluaran adalah median yang diminta, satu data per baris.
#### Soal3.go

```go
package main

import "fmt"

func median(arr [1000]int, n, data int){
	selesai:= false
	fmt.Print("Masukan Angka: ")
	for !selesai{
		fmt.Scan(&n)

		if n == -5313{
			selesai = true
		}else if n != 0{
			arr[data] = n
			data++
		}else if n  == 0{
			for i:=0; i<data-1; i++{
			min := i
			for j:=i+1; j<data; j++{
			if arr[min] > arr[j]{
				min = j
						}
					}
			temp:= arr[min]
			arr[min] = arr[i]
			arr[i] = temp
				}
				if data % 2 == 1{
					median:= data /2
					fmt.Println("Median: ", arr[median])
				}else{
					median:= (arr[data/2-1] + arr[data] / 2) / 2
					fmt.Println("Median: ", arr[median])
				}
			}
	}
}
func main() {
	var arr [1000]int
	var n, data int
	
	median(arr, n, data)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas12/output_tugas12/Soal3.png)
[penjelasan]
Program ini digunakan untuk mencari nilai median menggunakan bantuan sorting, program akan berhenti sata input -5313. Program akan melakukan sorting saat input = 0, jadi input yang tidak sama dengan 0 itu disimpan variabel arr[data], kalau input = 0 maka akan dilakukan sorting, sortingnya data yang mana bang? yang di sorting data yang diinputkan sebelum input = 0, nah setelah melakukan sorting lalu tentukan nilai mediannya dengan 2 cara, jika jumlah data ganjil. Median = jumlah data dibagi 2 maka ketemu nilai tengahnya tinggal di cetak nilai index nya, kalau genap beda karena nilai tengahnya ada 2 maka nilai idx kiri + nilai idx kanan lalu di bagi 2

