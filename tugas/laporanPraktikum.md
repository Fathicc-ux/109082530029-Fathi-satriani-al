# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Fathi satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Menukar nilai]
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (satu, dua, tiga string
	temp string)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas/output/soal1.png)
[penjelasan]
Program meminta input sebanyak 3 kali, yang dimana output awal hasilnya sama dengan input kesatu kedua dan ketiga
untuk output akhir berbeda dengan output awal yang dimana setiap inputannya diubah, contoh satu = dua, dua = tiga, tiga = temp dan  temp = 1. Inputan saya di program adalah (satu, dua, dan tiga) output awal = satu dua tiga, output akhir = dua tiga satu. Kenapa tiga = satu? padahal tiga = temp, karena itu sudah dirubah yang awalnya temp dirubah menjadi satu



### 2. [praktikum kimia]
Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
#### soal2.go

```go
package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var berhasil bool
	berhasil = true

	for i:= 1; i <= 5; i++{
		fmt.Scan(&w1, &w2, &w3, &w4)
		fmt.Println("Percobaan ", i, ":", w1, " ", w2, " ", w3, " ", w4, " ")
		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu"{
			berhasil = false
		}
	}
	fmt.Print("Berhasil: ", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas/output/soal2.png)
[penjelasan]
program digunakan untuk memberikan nilai kebenaran, Output akan bernilai TRUE saat kita input sesuai urutan warna pada setiap percobaan dan akan bernilai FALSE saat ada 1 percobaan yang input warnanya tidak sesuai urutan. yang dimana urutan warna nya adalah: merah kuning hijau ungu, bisa dilihat di output pada percobaan pertama bernilai TRUE dan FALSE pada percobaan kedua karena input pada baris pertama adalah: ungu hijau kuning merah, walaupun ke empat percobaan urut tetapi saat ada 1 percobaan yang warnannya tidak urut maka akan bernilai FALSE


### 3. [Aplikasi perhitungan biaya kirim berdasarkan berat parsel]
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
package main

import "fmt"

func main() {
	var parsel, biaya, Rp, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)

	kg := parsel / 1000
	detail_berat:= parsel % 1000
	fmt.Println("Detail berat: ", kg, "kg", "+" ,detail_berat, "gram")
	Rp = kg * 10000

	if parsel > 10000{
		biaya = 0
	}else  if detail_berat >= 500{
		biaya = 5 * detail_berat 
	}else {
		biaya = 15 * detail_berat
	}
	fmt.Println("Detail biaya: ", "Rp. ", Rp, " + ", "Rp. ", biaya)

	total = Rp + biaya

	fmt.Print("Total biaya: ","Rp. ", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas/output/soal3.png)
[penjelasan]
program digunakan untuk menghitung biaya kirim berdasarkan berat barang, biaya jasa 10.000/kg. Jika sisa berat tidak kurang dari 500 gr dikenakan biaya tambahan Rp.5/gr, tapi jika kurang dari 500gr biaya tambahan Rp.15/gr nya dan gratis apabila total berat lebih dari 10.000
contoh input: 8.500, sisa berat 500gr karena satuannya kg dan sisanya gr, karena sisa berat tidak kurang dari 500gr maka biaya tambahannya Rp.5/ gr =, 500 * Rp. 5 = 2.500. Total nya maka 80.000 + 2.500 =82.500