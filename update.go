package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const MAX = 100

type User struct {
	Username string
	Password string
}

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

var users []User
var dataResume [MAX]Resume
var jumlahData int
var pilihan int

func landingPage() {
	clearScreen()
	fmt.Println("-----------------------")
	fmt.Println(" Aplikasi Dummy ")
	fmt.Println("-----------------------")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("-----------------------")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		register()
	case 2:
		login()
	case 3:
		fmt.Println("Keluar Aplikasi, terima kasih sudah mampir hehehe")
	default:
		fmt.Println("Masukan tidak valid, mohon input sesuai nomor")
		landingPage()
	}
}

func Menu() {
	var pilihan string
	for pilihan != "8" {
		fmt.Println("\n========= MENU UTAMA =========")
		fmt.Println("1. Tambah Resume & Surat Lamaran")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Cari Data (Binary Search)")
		fmt.Println("4. Edit Data")
		fmt.Println("5. Hapus Data")
		fmt.Println("6. Urutkan Data Ascending (Insertion Sort)")
		fmt.Println("7. Urutkan Data Descending (Selection Sort)")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)
		clearScreen()

		switch pilihan {
		case "1":
			TambahResume()
		case "2":
			TampilData()
		case "3":
			CariData()
		case "4":
			EditData()
		case "5":
			HapusData()
		case "6":
			UrutkanDataAsc()
			fmt.Println("Data berhasil diurutkan secara ascending.")
		case "7":
			UrutkanDataDesc()
			fmt.Println("Data berhasil diurutkan secara descending.")
		case "8":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}

func register() {
	clearScreen()
	var username, password string

	fmt.Println("-----------------------")
	fmt.Println(" Register Dummy ")
	fmt.Println("-----------------------")
	fmt.Print(" Username: ")
	fmt.Scanln(&username)
	fmt.Print(" Password: ")
	fmt.Scanln(&password)

	for _, user := range users {
		if user.Username == username {
			fmt.Println("\n   username telah digunakan")
			landingPage()
			return
		}
	}
	users = append(users, User{Username: username, Password: password})
	fmt.Println("\n   Akun berhasil didaftarkan")
	landingPage()
}

func login() {
	var inputUsn, inputPass string
	clearScreen()
	fmt.Println("-----------------------")
	fmt.Println(" Login Dummy ")
	fmt.Println("-----------------------")
	fmt.Print(" Username: ")
	fmt.Scanln(&inputUsn)
	fmt.Print(" Password: ")
	fmt.Scanln(&inputPass)

	clearScreen()
	for _, user := range users {
		if user.Username == inputUsn && user.Password == inputPass {
			fmt.Println("\n   ---- Selamat Datang", user.Username, "-----")
			Menu()
			return
		}
	}
	fmt.Println("\n Username atau Password salah!")
	landingPage()
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

func TambahResume() {
	if jumlahData >= MAX {
		fmt.Println("Data penuh, tidak bisa menambah resume baru.")
		return
	}

	var r Resume
	r.ID = jumlahData + 1

	fmt.Print("Nama: ")
	fmt.Scanln(&r.Nama)
	fmt.Print("Email: ")
	fmt.Scanln(&r.Email)
	fmt.Print("Nomor HP: ")
	fmt.Scanln(&r.NomorHP)
	fmt.Print("Pendidikan: ")
	fmt.Scanln(&r.Pendidikan)
	fmt.Print("Pengalaman: ")
	fmt.Scanln(&r.Pengalaman)
	fmt.Print("Keahlian: ")
	fmt.Scanln(&r.Keahlian)
	fmt.Print("Posisi yang dilamar: ")
	var posisi string
	fmt.Scanln(&posisi)

	r.SuratLamaran = GenerateSuratTemplate(r.Nama, posisi, r.Keahlian)

	dataResume[jumlahData] = r
	jumlahData++
	fmt.Println("Resume berhasil dibuat!")
}

func TampilData() {
	clearScreen()
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

func BinarySearch(nama string) int {
	clearScreen()
	low, high := 0, jumlahData-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(dataResume[mid].Nama, nama) {
			return mid
		} else if strings.ToLower(nama) < strings.ToLower(dataResume[mid].Nama) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func CariData() {
	clearScreen()
	UrutkanDataAsc()
	var nama string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&nama)
	indeks := BinarySearch(nama)

	if indeks != -1 {
		r := dataResume[indeks]
		fmt.Printf("\nID: %d\nNama: %s\nEmail: %s\nNo HP: %s\nPendidikan: %s\nPengalaman: %s\nKeahlian: %s\nSurat Lamaran:\n%s\n",
			r.ID, r.Nama, r.Email, r.NomorHP, r.Pendidikan, r.Pengalaman, r.Keahlian, r.SuratLamaran)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func EditData() {
	clearScreen()
	var id int
	fmt.Print("Masukkan ID yang ingin diedit: ")
	fmt.Scanln(&id)

	indeks := -1
	for i := 0; i < jumlahData; i++ {
		if dataResume[i].ID == id {
			indeks = i
			break
		}
	}

	if indeks == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru: ")
	fmt.Scanln(&dataResume[indeks].Nama)
	fmt.Print("Email baru: ")
	fmt.Scanln(&dataResume[indeks].Email)
	fmt.Print("Nomor HP baru: ")
	fmt.Scanln(&dataResume[indeks].NomorHP)
	fmt.Print("Pendidikan baru: ")
	fmt.Scanln(&dataResume[indeks].Pendidikan)
	fmt.Print("Pengalaman baru: ")
	fmt.Scanln(&dataResume[indeks].Pengalaman)
	fmt.Print("Keahlian baru: ")
	fmt.Scanln(&dataResume[indeks].Keahlian)
	fmt.Print("Posisi yang dilamar: ")
	var posisi string
	fmt.Scanln(&posisi)

	dataResume[indeks].SuratLamaran = GenerateSuratTemplate(dataResume[indeks].Nama, posisi, dataResume[indeks].Keahlian)
	fmt.Println("Data berhasil diedit.")
}

func HapusData() {
	clearScreen()
	var id int
	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scanln(&id)

	indeks := -1
	for i := 0; i < jumlahData; i++ {
		if dataResume[i].ID == id {
			indeks = i
			break
		}
	}

	if indeks == -1 {
		fmt.Println("ID tidak ditemukan.")
		return
	}

	for i := indeks; i < jumlahData-1; i++ {
		dataResume[i] = dataResume[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus.")
}

func UrutkanDataAsc() {
	clearScreen()
	for i := 1; i < jumlahData; i++ {
		temp := dataResume[i]
		j := i - 1
		for j >= 0 && strings.ToLower(dataResume[j].Nama) > strings.ToLower(temp.Nama) {
			dataResume[j+1] = dataResume[j]
			j--
		}
		dataResume[j+1] = temp
	}
}

func UrutkanDataDesc() {
	clearScreen()
	for i := 0; i < jumlahData-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if strings.ToLower(dataResume[j].Nama) > strings.ToLower(dataResume[maxIdx].Nama) {
				maxIdx = j
			}
		}
		if maxIdx != i {
			dataResume[i], dataResume[maxIdx] = dataResume[maxIdx], dataResume[i]
		}
	}
}

func main() {
	landingPage()
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Platform tidak didukung.")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
