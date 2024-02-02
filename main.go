package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	b := 70
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
