package main

import "fmt"

func main () {
	var nama, univ string
	var nim int


	fmt.Print("masukan nama anda : ")
	fmt.Scan(&nama)
	fmt.Print("masukan nim anda : ")
	fmt.Scan(&nim)
	fmt.Print("masukan univ anda : ")
	fmt.Scan(&univ)

	fmt.Printf("Perkenalkan nama saya %v dengan nim %d dan saya sedang berkuliah di %v" ,nama, nim, univ)
	

}