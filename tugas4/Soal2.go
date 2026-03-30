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