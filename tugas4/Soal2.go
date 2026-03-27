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
  var nama1, nama2 string
  var soal1, soal2, skor1, skor2 int
  
  fmt.Scan(&nama1)
  hitungskor(&soal1, &skor1)
  
  fmt.Scan(&nama2)
  hitungskor(&soal2, &skor2)
  
  if soal1 > soal2{
      fmt.Println(nama1, soal1, skor1)
  }else if soal2 > soal1{
      fmt.Println(nama2, soal2, skor2)
  }else if skor1 < skor2{
      fmt.Println(nama1, soal1, skor1)
  }else{
      fmt.Println(nama2, soal2, skor2)
  }
}