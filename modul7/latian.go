package main

import ("fmt")

type rek struct{
	norek int
	pin int
	saldo int
	tarik int
}
func main () {
	var atm rek
	atm.norek = 889000
	atm.pin = 123456
	atm.saldo = 50000
	fmt.Print("Masukan nomor rekning: ")
	fmt.Scan(&atm.norek)
	fmt.Print("Masukan pin: ")
	fmt.Scan(&atm.pin)

	if atm.norek == 889000 && atm.pin == 123456{
		fmt.Print("Mau tarik berapa?: ")
			fmt.Scan(&atm.tarik)
		atm.saldo = atm.saldo - atm.tarik
			if atm.saldo < atm.tarik {
		fmt.Print("Saldo tidak memenuhi")
			}else{
		fmt.Print("Sisa saldo anda sekarang adalah Rp. ", atm.saldo)
			}
	}else if atm.norek == 889000 || atm.pin != 123456{
		fmt.Print("Pin anda salah")
	}else if atm.norek != 889000 || atm.pin == 123456{
		fmt.Print("Nomor rekening tidak terdaftar")
	}
}