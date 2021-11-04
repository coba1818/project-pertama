package main

import (
	"fmt"
)

func main() {
	var angka1 int
	var angka2 int
	var jumlah int
	{

		fmt.Print("Masukkan Angka Pertama : ")
		fmt.Scanln(&angka1)
		fmt.Print("Masukkan Angka Pertama : ")
		fmt.Scanln(&angka2)
		jumlah = angka1 + angka2
		fmt.Print("Total dari ", angka1, " + ", angka2, " adalah ", jumlah)

	}
}
