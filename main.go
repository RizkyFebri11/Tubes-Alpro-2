package main

import "fmt"

type Resume struct {
	ID           int
	Nama         string
	Email        string
	NomorHP      string
	Pendidikan   string
	Pengalaman   string
	Keahlian     string
	SuratLamaran string
}

func GenerateSuratTemplate(nama, posisi, keahlian string) string {
	surat := fmt.Sprintf(`Yth. HRD %s,

Dengan hormat,

Saya yang bertanda tangan di bawah ini:
%s

Dengan ini mengajukan lamaran kerja untuk posisi %s. Saya memiliki keahlian dalam %s dan yakin dapat memberikan kontribusi positif bagi perusahaan Anda.

Saya berharap dapat diberikan kesempatan untuk mengikuti proses seleksi lebih lanjut. Atas perhatian dan kesempatannya, saya ucapkan terima kasih.

Hormat saya,
%s`, posisi, nama, posisi, keahlian, nama)

	return surat
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("    AI RESUME & LAPORAN    ")
	fmt.Println("-----------------------")
	fmt.Println("1. Data Pengguna")
	fmt.Println("2. Buat Resume")
	fmt.Println("3. Buat Surat Lamaran")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")

}

func InputDataPengguna() {

}

func BuatResume() {

}

func BuatSuratLamaran() {

}

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			InputDataPengguna()
		case 2:
			BuatResume()
		case 3:
			BuatSuratLamaran()
		case 4:
			fmt.Println("Program berhenti, sampai jumpa lain waktu!")
		default:
			fmt.Println("Pilihan Tidak Valid, SIlakan Coba Lagi.")
		}
		fmt.Println()
	}
}

