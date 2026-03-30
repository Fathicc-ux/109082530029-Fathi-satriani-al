# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[Fathi Satriani Al jauzy] - [109082530029]</p>

## Unguided 

### 1. [Faktorial, Permutasi, Kombinasi dengan prosedur]
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
#### Soal1.go

```go
package main

import "fmt"

func faktorial(n int, hasil* int){
	*hasil = 1
	for i:=1; i<=n; i++{
		*hasil = *hasil * i
	}
}
func permutasi(n, r int, hasil* int) {
    var nf, fn int
    faktorial(n, &nf)
    faktorial(n-r, &fn)
	*hasil = nf/fn
}
func kombinasi(n, r int, hasil* int) {
    var nf, nr, fn int
    faktorial(n, &nf)
    faktorial(r, &nr)
    faktorial(n-r, &fn)
	*hasil = nf / (nr * fn)
}
func main () {
	var a, b, c, d int
	var per1, per2, kom1, kom2 int

	fmt.Scan(&a, &b, &c, &d)
	
	permutasi(a, c, &per1)
	kombinasi(a, c, &kom1)
	permutasi(b, d, &per2)
	kombinasi(b, d, &kom2)
	
	fmt.Println("Baris 1: ", per1, kom1)
	fmt.Println("Baris 2: ", per2, kom2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas4/output_tugas4/Soal1.png)
[penjelasan]
Program ini dibuat untuk menghitung faktorial, permutasi, dan kombinasi menggunakan prosedur.
Di dalam program ini terdapat tiga prosedur, yaitu faktorial, permutasi, dan kombinasi.
pada prosedur faktorial, terdapat input n dan parameter hasil. Parameter hasil digunakan untuk menyimpan nilai faktorial dari n. prosedur permutasi, terdapat input n dan r, serta parameter hasil. Di sini digunakan variabel tambahan yaitu nf dan fn. Variabel nf digunakan untuk menyimpan nilai n!, sedangkan fn untuk menyimpan (n-r)!. prosedur kombinasi, terdapat input n dan r, serta parameter hasil. Di sini saya menggunakan variabel tambahan yaitu nf, nr dan fn. Variabel nf digunakan untuk menyimpan nilai n!, nr untuk menyimpan nilai r!, sedangkan fn untuk menyimpan (n-r)!.
program akan menginput 4 angka a, b, c, dan d. Kemudian dibuat variabel baru untuk hasil permutasi dan kombinasi. Panggil prosedur dan simpan hasilnya ke variabel baru yang sudah dibuat, setelah itu tinggal print variabel baru yang digunakan untuk menyimpan prosedur permutasi dan kombinasi

### 2. [Program untuk mencari pemenang dengan prosedur]
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil
diselesaikan.
#### Soal2.go

```go
package main

import "fmt"

func hitungskor(soal, skor *int){
    var waktu int

    *soal = 0
    *skor = 0

    for i := 0; i < 8; i++ {
        fmt.Scan(&waktu)

        if waktu <= 300 {
            *soal = *soal + 1
            *skor = *skor + waktu
        }
    }
}
func main() {
  var nama, pemenang string
  var soal1, soal2, skor1, skor2 int

  for selesai := false; !selesai;{
    fmt.Scan(&nama)
  if nama == "selesai"{
        selesai = true
    }else{
    hitungskor(&soal2, &skor2)
    }
    if skor2 > skor1 || (soal2 == soal1 && skor2 < skor1){
        pemenang = nama
        soal1 = soal2
        skor1 = skor2
    }else if soal2 > soal1{
        pemenang = nama
        soal1 = soal2
        skor1 =  skor2
        }
    }
    fmt.Println(pemenang, soal1, skor1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Fathicc-ux/109082530029-Fathi-satriani-al/blob/main/tugas4/output_tugas4/Soal2.png)
[penjelasan]
Program ini dibuat untuk mencari pemenang dari suatu kompetisi, fungsi hitungskor digunakan untuk menghitung jumlah soal yang berhasil dikerjakan dan total waktu (skor) yang diperoleh. Aturan Penentuan pemenang: 1. yang mengerjakan soal lebih banyak akan menjadi pemenang. 2. Jika jumlah soal yang dikerjakan sama, pemenang ditentukan dari waktu yang tercepat (skor yang lebih kecil). Program akan terus berjalan (looping) dan berhenti saat input nama == "selesai", program akan terus membandingkan jumlah soal dan skor peserta yang telah di inputkan jika ada yang lebih sedikit waktu (skor) dan jumlah soalnya sesuai aturan penentuan pemenang maka peserta tersebut yang akan menjadi pemenang. Program akan mencetak nama pemenang, jumlah pemenang dan banyak skor saat input nama == "selesai"
