//Fathi satriani al _ 109082530029
package main

import "fmt"

const nMAX int = 1000

type lulusan struct{
	nama, nim string
	semester int
	eprt, ipk float64
}
type tabLulus [nMAX] lulusan

func isidata(t *tabLulus, n *int){
	var nim string
	*n = 0
	fmt.Scan(&nim)
	for nim != "none" && *n < nMAX{
		t[*n].nim = nim
		fmt.Scan(&t[*n].nama, &t[*n].eprt, &t[*n].semester, &t[*n].ipk)
		*n = *n + 1
		fmt.Scan(&nim)
	}
}
func MAXeprt(t *tabLulus, n *int) float64{
	var max float64
	var i int

	max = t[0].eprt
	i = 1
	for i < *n {
		if max < t[i].eprt{
			max = t[i].eprt
		}
		i = i + 1
	}
	return max
}
func minIpK(t *tabLulus, n *int) float64{
	var min lulusan
	var i int

	min = t[0]
	i = 1
	for i < *n{
		if min.ipk > t[i].ipk{
			min = t[i]
		}
		i = i + 1
	}
	return min.ipk
}
func rataan(t *tabLulus, n *int) float64{
	var jum float64
	
	jum = 0
	
	for i := 0; i <= *n; i++{
		jum = jum + float64(t[i].semester)
	}
	return jum / float64(*n)
}
func main() {
	var dataMHS tabLulus
	var nMHS int
	var eprtTertinggi, ipkTerendah, rerataSemester float64

	isidata(&dataMHS, &nMHS)
	eprtTertinggi = MAXeprt(&dataMHS, &nMHS)
	ipkTerendah = minIpK(&dataMHS, &nMHS)
	rerataSemester = rataan(&dataMHS, &nMHS)
	fmt.Print(eprtTertinggi, ipkTerendah,rerataSemester )
}