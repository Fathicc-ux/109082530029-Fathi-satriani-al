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
fatih jelek banget



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
fatih jelek banget


tadi done, untuk mengganti foto ss masing masing soal tinggal ss masukan ke folder output kemudian ganti nama, pastikan pada nama file fotonya sama