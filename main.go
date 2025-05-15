package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX = 100

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

var dataResume [MAX]Resume
var jumlahData int

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

func TambahResume() {
	if jumlahData >= MAX {
		fmt.Println("Data penuh, tidak bisa menambah resume baru.")
		return
	}

	var r Resume
	r.ID = jumlahData + 1

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama: ")
	r.Nama, _ = reader.ReadString('\n')
	r.Nama = strings.TrimSpace(r.Nama)

	fmt.Print("Email: ")
	r.Email, _ = reader.ReadString('\n')
	r.Email = strings.TrimSpace(r.Email)

	fmt.Print("Nomor HP: ")
	r.NomorHP, _ = reader.ReadString('\n')
	r.NomorHP = strings.TrimSpace(r.NomorHP)

	fmt.Print("Pendidikan: ")
	r.Pendidikan, _ = reader.ReadString('\n')
	r.Pendidikan = strings.TrimSpace(r.Pendidikan)

	fmt.Print("Pengalaman: ")
	r.Pengalaman, _ = reader.ReadString('\n')
	r.Pengalaman = strings.TrimSpace(r.Pengalaman)

	fmt.Print("Keahlian: ")
	r.Keahlian, _ = reader.ReadString('\n')
	r.Keahlian = strings.TrimSpace(r.Keahlian)

	fmt.Print("Posisi yang dilamar: ")
	posisi, _ := reader.ReadString('\n')
	posisi = strings.TrimSpace(posisi)

	r.SuratLamaran = GenerateSuratTemplate(r.Nama, posisi, r.Keahlian)

	dataResume[jumlahData] = r
	jumlahData++
	fmt.Println("Resume berhasil dibuat!")
}

func TampilData() {
	fmt.Println("\n======= DAFTAR RESUME =======")
	for i := 0; i < jumlahData; i++ {
		r := dataResume[i]
		fmt.Printf("ID: %d\nNama: %s\nEmail: %s\nNo HP: %s\nPendidikan: %s\nPengalaman: %s\nKeahlian: %s\nSurat Lamaran:\n%s\n-----------------------------\n",
			r.ID, r.Nama, r.Email, r.NomorHP, r.Pendidikan, r.Pengalaman, r.Keahlian, r.SuratLamaran)
	}
	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println(" AI RESUME & LAPORAN ")
	fmt.Println("-----------------------")
	fmt.Println("1. Data Pengguna")
	fmt.Println("2. Tampilkan Resume & Surat Lamaran")
	fmt.Println("3. Exit")
	fmt.Println("-----------------------")

}

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			TambahResume()
		case 2:
			TampilData()
		case 3:
			fmt.Println("Program berhenti, semoga sukses dengan lamarannya!")
		default:
			fmt.Println("Pilihan Tidak Valid, SIlakan Coba Lagi.")
		}
		fmt.Println()
		if pilih == 3 {
			break
		}
	}
}
