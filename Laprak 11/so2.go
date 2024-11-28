package main

import "fmt"

func main() {
	var tipe_kendaraan string
	var durasi, tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&tipe_kendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch {
	case tipe_kendaraan == "Motor" && durasi < 1 && durasi >= 1:
		tarif = 2000
	case tipe_kendaraan == "Motor" && durasi > 1:
		tarif = 2000 * durasi
	case tipe_kendaraan == "Mobil" && durasi < 1 && durasi >= 1:
		tarif = 5000
	case tipe_kendaraan == "Mobil" && durasi > 1:
		tarif = 5000 * durasi
	case tipe_kendaraan == "Truk" && durasi < 1 && durasi >= 1:
		tarif = 8000
	case tipe_kendaraan == "Truk" && durasi > 1:
		tarif = 8000 * durasi

	default:

		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif parkir : Rp %d", tarif)
}
