package main

import "fmt"

func main() {
	var bilangan, hasil int
	fmt.Scan(&bilangan)

	switch {
	case bilangan%2 != 0:
		hasil = bilangan + (bilangan + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)

	case bilangan%5 == 0 && bilangan%10 != 0:
		hasil = bilangan * bilangan
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, hasil)

	case bilangan%10 == 0:
		hasil = bilangan / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)

	default:
		hasil = bilangan * (bilangan + 1)

		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
	}

}
