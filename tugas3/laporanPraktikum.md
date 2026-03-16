# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Fathi satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Menukar nilai]
Telusuri 
#### Soal1.go

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
Program meminta input sebanyak 3 kali, yang dimana output awal hasilnya sama dengan input kesatu 


### 2. [praktikum kimia]
Di setiap 
#### Soal2.go

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
program 


### 3. [Aplikasi perhitungan biaya kirim berdasarkan berat parsel]
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan

#### Soal3.go

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
program 