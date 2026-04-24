package main
import "fmt"
type waktu struct {
	jam, menit, detik int
}
type rental struct{
	nama, naga string
	tagihan, biaya int
}
func main(){
var ren rental
var wParkir, wPulang, durasi waktu
var dParkir, dPulang, lParkir int

fmt.Print("Nama penyewa: ")
fmt.Scan(&ren.nama)
fmt.Print("Nama naga: ")
fmt.Scan(&ren.naga)

fmt.Print("waktu mulai: ")
fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
fmt.Print("Waktu selesai: ")
fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)


dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
lParkir = dPulang - dParkir
durasi.jam = lParkir / 3600
durasi.menit = lParkir % 3600 / 60
durasi.detik = lParkir % 3600 % 60

if ren.naga == "gembul"{
	ren.biaya = lParkir * 5000
}else{
	ren.biaya =  lParkir * 200
}
if lParkir > 3600{
	ren.tagihan += 15000
}
fmt.Println("tagihan: ", ren.tagihan)
fmt.Print("yth nama peneywa: ", ren.nama, " jenis naga: ",ren.naga, " durasi sewa: ",durasi.jam," jam ", durasi.menit," menit ",durasi.detik, " detik",
)
}