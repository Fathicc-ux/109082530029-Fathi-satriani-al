package main
import "fmt"

type data struct {
	nama string
  nilai float64
}
type siswa [100] data
type remidial [100] string

func hitungRata(S siswa) {
/*  I.S data sejumlah nama dan niai siswa telah siap pada piranti masukan
    proses: simpan data mahasiswa yang remedial pada array remidial dan hitung banyaknya mahasiswa yang remedial
    F.S array siswa dan remidial telah terisi sejumlah data dan menampilkan nilai rata-rata, banyaknya siswa yang remedial, serta nama-nama mahasiswa yang remedial*/
	
var R remidial
	var total, rata float64
	var n, jumlahRemidi int

	for n < 100 {
		fmt.Scan(&S[n].nama)

		if S[n].nama == "0" {
			break
		}

		fmt.Scan(&S[n].nilai)

		total += S[n].nilai

		// cek remedial
		if S[n].nilai < 50 {
			R[jumlahRemidi] = S[n].nama
			jumlahRemidi++
		}

		n++
	}

	rata = total / float64(n)
	fmt.Printf("%.2f\n", rata)
	fmt.Println(jumlahRemidi)

	for i := 0; i < jumlahRemidi; i++ {
		fmt.Println(R[i])
	}
}

func main() {
	var S siswa

	hitungRata(S)
}