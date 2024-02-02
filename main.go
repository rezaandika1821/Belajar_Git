package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	b := 70

	var sandi int
	fmt.Print("Masukkan Sandi Anda : ")
	fmt.Scanln(&sandi)
	if sandi != 230104 {
		fmt.Println("Sandi Salah")
		fmt.Println("Anda di Blokir")
		goto a
	} else if sandi == 230104 {

		fmt.Println("Hello World")
		fmt.Printf("ini umur %d kakek saya ", b)
		fmt.Println("Hello World")
		fmt.Print("Masukkan Umur Anda")
		fmt.Scan(&a)
		fmt.Println("ini umur saya ", a)

		fmt.Println(strings.Repeat("=", 20))
		fmt.Println()
		fmt.Println(strings.Repeat("=", 20))

	}
a:

//ini baru edit
}
