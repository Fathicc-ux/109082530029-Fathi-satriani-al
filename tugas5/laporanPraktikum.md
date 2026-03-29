# <h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Deret fibonacci]
Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.
#### Soal1.go

```go
package main
import "fmt"

func fibo(n int) int {
    if n == 0{
        return 0
    }else if n == 1{
        return 1
    }else{
    return fibo(n-1) + fibo(n-2)
	}
}
func main() {
    var n int
	fmt.Print("n: ")
    fmt.Scan(&n)

    for i:= 0; i<=n; i++{
    fmt.Print(i, " ")
	fmt.Println(fibo(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas5/output_tugas5/Soal1.png)
[penjelasan]
Program ini digunakan untuk menampilkan deret Fibonacci.Di mana setiap suku dihasilkan dari penjumlahan dua suku sebelumnya. Kalau n == 0 maka hasilnya 0, kalau n == 1 maka hasilnya 1, dan kalau bukan maka hitung fibo(n-1) + fibo(n-2). Input n lalu masuk ke perulangan, mulai perulangan dari 0 sampai n setiap perulangan terdapat hasil dari setiap suku.

### 2. [Pola bintang]
Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.
#### Soal2.go

```go
package main
import "fmt"

func bintang(n int){
    if n == 0{
        return
    }
    bintang(n-1)
    for i:=1; i<=n; i++{
        fmt.Print("*")
    }
    fmt.Println( )
}
func main() {
    var n int
    fmt.Print("Angka: ")
    fmt.Scan(&n)
    
    bintang(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas5/output_tugas5/Soal2.png)
[penjelasan]
Program ini digunakan untuk menampilkan pola bintang menggunakan fungsi rekursif. Fungsi bintang akan berhenti jika nilai n == 0, fungsi akan terus berjalan dan berhenti saat n == 0 dengan n-1. Contoh input 5 akan berjalan sampai 1 dan berhenti di 0 karen n == 0, Lalu masuk ke perulangan untuk mencetak bintang dimulai dari bintang 1 sampai bintang ke-n


### 3. [Mencari faktor dengan rekrusif]
Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### Soal3.go

```go
package main
import "fmt"

func faktor(n, i int) {
    if i > n{
        return
    }
    if n%i == 0{
        fmt.Print(i, " ")
    }
    faktor(n, i+1)
}
func main() {
    var n int
    fmt.Scan(&n)

    faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas5/output_tugas5/Soal3.png)
[penjelasan]
Program ini digunakan untuk mencari faktor dari suatu bilangan n menggunakan rekursif. Program akan berhenti jika i > n, Nilai i didapat dari fungsi main, dimulai dari 1 faktor(n, 1). Program akan mengecek apakah n % i == 0, jika hasilnya 0 berarti i adalah faktor dari n. Fungsi akan berjalan terus dengan i+1, yang digunakan untuk mengcek faktor dari n dan akan berhenti saat nilai i lebih dari n