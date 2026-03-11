# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Fathi satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Soal]
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



### 2. [Soal]
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


### 2. [Soal]
#### soal3.go

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
program digunakan untuk memberikan nilai kebenaran, Output akan bernilai TRUE saat kita input sesuai urutan warna pada setiap