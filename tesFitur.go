// Kode dibuat oleh Rizky Febriyanto dan Orlando Silas Davinci Kambu IF-48-07

/*
"math/rand" & "time" nanti dia buat random pick templatenya gituu
"os/exec" & "runtime" sebagai fungsi dari clearscreen
"strings" buat struktur sebuah string yang terikat dengan array (digunakan untuk search dan sort)
*/

package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const NMAX = 100

type User struct {
	Username string
	Password string
}

type Resume struct {
	ID            int
	Nama          string
	Email         string
	NomorHP       string
	Pendidikan    string
	Pengalaman    string
	Keahlian      string
	PosisiDilamar string
	SuratLamaran  string
}

var users []User
var dataResume []Resume

func main() {
	rand.Seed(time.Now().UnixNano())
	landingPage()
}

// landing page awal
func landingPage() {
	clearScreen()
	fmt.Println("===================================")
	fmt.Println(" Aplikasi Pembuat Resume & Lamaran ")
	fmt.Println("===================================")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("===================================")
	fmt.Print("Pilih opsi: ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka (1-3).")
		pressEnterToContinue()
		landingPage()
		return
	}

	switch choice {
	case 1:
		register()
	case 2:
		login()
	case 3:
		fmt.Println("Terima kasih sudah menggunakan aplikasi ini!")
	default:
		fmt.Println("Pilihan tidak ada. Silakan masukkan angka 1, 2, atau 3.")
		pressEnterToContinue()
		landingPage()
	}
}

// main menu pengguna yang nantinya akan menjadi menu utama sebuah program
func mainMenu() {
	for {
		clearScreen()
		fmt.Println("\n========= MENU UTAMA =========")
		fmt.Println("1. Tambah Resume & Surat Lamaran")
		fmt.Println("2. Tampilkan Semua Resume")
		fmt.Println("3. Cari Resume (Berdasarkan Nama)")
		fmt.Println("4. Edit Data Resume")
		fmt.Println("5. Hapus Data Resume")
		fmt.Println("6. Urutkan Resume Menaik (A-Z)")
		fmt.Println("7. Urutkan Resume Menurun (Z-A)")
		fmt.Println("8. Keluar ke Menu Awal")
		fmt.Print("Pilih menu: ")

		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid. Silakan coba lagi.")
			pressEnterToContinue()
			continue
		}

		switch choice {
		case "1":
			tambahResume()
		case "2":
			tampilkanSemuaResume()
		case "3":
			cariResume()
		case "4":
			editResume()
		case "5":
			hapusResume()
		case "6":
			urutkanResumeAscending()
			fmt.Println("\nData resume berhasil diurutkan secara menaik (A-Z) berdasarkan nama.")
		case "7":
			urutkanResumeDescending()
			fmt.Println("\nData resume berhasil diurutkan secara menurun (Z-A) berdasarkan nama.")
		case "8":
			fmt.Println("Anda telah keluar dari menu utama.")
			pressEnterToContinue()
			landingPage()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		pressEnterToContinue()
	}
}

// fungsi register pengguna untuk tahap awal aplikasi
func register() {
	clearScreen()
	var username, password string

	fmt.Println("=======================")
	fmt.Println(" Daftar Akun Baru ")
	fmt.Println("=======================")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)
	if username == "" {
		fmt.Println("Username tidak boleh kosong.")
		pressEnterToContinue()
		register()
		return
	}

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)
	if password == "" {
		fmt.Println("Password tidak boleh kosong.")
		pressEnterToContinue()
		register()
		return
	}

	for _, user := range users {
		if user.Username == username {
			fmt.Println("\nUsername ini sudah dipakai. Silakan pilih username lain.")
			pressEnterToContinue()
			landingPage()
			return
		}
	}

	users = append(users, User{Username: username, Password: password})
	fmt.Println("\nAkun berhasil didaftarkan! Silakan login.")
	pressEnterToContinue()
	landingPage()
}

// Login Pengguna
func login() {
	var inputUsn, inputPass string
	clearScreen()
	fmt.Println("=======================")
	fmt.Println(" Masuk Aplikasi ")
	fmt.Println("=======================")
	fmt.Print("Username: ")
	fmt.Scanln(&inputUsn)
	fmt.Print("Password: ")
	fmt.Scanln(&inputPass)

	clearScreen()
	for _, user := range users {
		if user.Username == inputUsn && user.Password == inputPass {
			fmt.Printf("\n---- Selamat Datang, %s! -----\n", user.Username)
			pressEnterToContinue()
			mainMenu()
			return
		}
	}
	fmt.Println("\nUsername atau Password salah. Silakan coba lagi.")
	pressEnterToContinue()
	landingPage()
}

func tambahResume() {
	if len(dataResume) >= NMAX { // Diubah dari MAX_RESUMES ke NMAX
		fmt.Println("Maaf, penyimpanan resume sudah penuh. Anda tidak bisa menambah data baru.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var r Resume
	r.ID = len(dataResume) + 1

	fmt.Println("\n--- Buat Resume Baru ---")
	fmt.Print("Nama Lengkap: ")
	namaLengkap, _ := reader.ReadString('\n')
	r.Nama = strings.TrimSpace(namaLengkap)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	r.Email = strings.TrimSpace(email)

	fmt.Print("Nomor HP: ")
	nomorHP, _ := reader.ReadString('\n')
	r.NomorHP = strings.TrimSpace(nomorHP)

	fmt.Print("Pendidikan Terakhir: ")
	pendidikan, _ := reader.ReadString('\n')
	r.Pendidikan = strings.TrimSpace(pendidikan)

	fmt.Print("Pengalaman Kerja (pisahkan dengan koma jika lebih dari satu): ")
	pengalaman, _ := reader.ReadString('\n')
	r.Pengalaman = strings.TrimSpace(pengalaman)

	fmt.Print("Keahlian (pisahkan dengan koma jika lebih dari satu): ")
	keahlian, _ := reader.ReadString('\n')
	r.Keahlian = strings.TrimSpace(keahlian)

	fmt.Print("Posisi yang dilamar: ")
	posisiDilamar, _ := reader.ReadString('\n')
	r.PosisiDilamar = strings.TrimSpace(posisiDilamar)

	fmt.Print("Nama Perusahaan yang dilamar: ")
	namaPerusahaan, _ := reader.ReadString('\n')
	namaPerusahaan = strings.TrimSpace(namaPerusahaan)

	r.SuratLamaran = generateSuratTemplate(namaPerusahaan, r.Nama, r.PosisiDilamar, r.Keahlian)

	dataResume = append(dataResume, r)
	fmt.Println("\nResume berhasil dibuat dan surat lamaran otomatis tergenerasi!")
}

// menampilkan data hasil inputan dari "TambahResume" yang nanti akan digenerate oleh "GenerateSuratTemplate"
func tampilkanSemuaResume() {
	clearScreen()
	fmt.Println("\n======= DAFTAR SEMUA RESUME =======")
	if len(dataResume) == 0 {
		fmt.Println("Belum ada data resume yang tersimpan.")
		return
	}

	for _, r := range dataResume {
		fmt.Printf("ID: %d\n", r.ID)
		fmt.Printf("Nama: %s\n", r.Nama)
		fmt.Printf("Email: %s\n", r.Email)
		fmt.Printf("No HP: %s\n", r.NomorHP)
		fmt.Printf("Pendidikan: %s\n", r.Pendidikan)
		fmt.Printf("Pengalaman: %s\n", r.Pengalaman)
		fmt.Printf("Keahlian: %s\n", r.Keahlian)
		fmt.Printf("Posisi Dilamar: %s\n", r.PosisiDilamar)
		fmt.Printf("Surat Lamaran:\n%s\n", r.SuratLamaran)
		fmt.Println("======================================")
	}
}

// fungsi pencarian menggunakan model Binary Search
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

// mencari data dari "dataResume" dengan string yang tersimpan
func cariResume() {
	if len(dataResume) == 0 {
		fmt.Println("Belum ada data untuk dicari.")
		return
	}

	urutkanResumeAscending()

	clearScreen()
	var namaCari string
	fmt.Print("Masukkan nama yang dicari: ")
	reader := bufio.NewReader(os.Stdin)
	namaCariInput, _ := reader.ReadString('\n')
	namaCari = strings.TrimSpace(namaCariInput)

	indeks := binarySearch(namaCari)

	if indeks != -1 {
		r := dataResume[indeks]
		fmt.Println("\n--- Data Ditemukan ---")
		fmt.Printf("ID: %d\nNama: %s\nEmail: %s\nNo HP: %s\nPendidikan: %s\nPengalaman: %s\nKeahlian: %s\nPosisi Dilamar: %s\nSurat Lamaran:\n%s\n",
			r.ID, r.Nama, r.Email, r.NomorHP, r.Pendidikan, r.Pengalaman, r.Keahlian, r.PosisiDilamar, r.SuratLamaran)
	} else {
		fmt.Println("Data tidak ditemukan. Pastikan nama yang Anda masukkan sudah benar.")
	}
}

// mengedit data "TampilData" yang dibuat melalui "GenerateSuratTemplate" melalui "ID"
func editResume() {
	clearScreen()
	if len(dataResume) == 0 {
		fmt.Println("Belum ada data resume yang tersimpan untuk diedit.")
		return
	}

	var id int
	fmt.Print("Masukkan ID resume yang ingin diedit: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input ID tidak valid. Harap masukkan angka.")
		return
	}
	indeks := -1
	for i, r := range dataResume {
		if r.ID == id {
			indeks = i
			break
		}
	}
	if indeks == -1 {
		fmt.Println("ID resume tidak ditemukan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var input string

	fmt.Println("\n--- Edit Data Resume ---")
	fmt.Printf("Nama lama: %s\n", dataResume[indeks].Nama)
	fmt.Print("Nama baru (kosongkan jika tidak diubah): ")
	inputBytes, _ := reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].Nama = input
	}

	fmt.Printf("Email lama: %s\n", dataResume[indeks].Email)
	fmt.Print("Email baru (kosongkan jika tidak diubah): ")
	inputBytes, _ = reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].Email = input
	}

	fmt.Printf("Nomor HP lama: %s\n", dataResume[indeks].NomorHP)
	fmt.Print("Nomor HP baru (kosongkan jika tidak diubah): ")
	inputBytes, _ = reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].NomorHP = input
	}

	fmt.Printf("Pendidikan lama: %s\n", dataResume[indeks].Pendidikan)
	fmt.Print("Pendidikan baru (kosongkan jika tidak diubah): ")
	inputBytes, _ = reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].Pendidikan = input
	}

	fmt.Printf("Pengalaman lama: %s\n", dataResume[indeks].Pengalaman)
	fmt.Print("Pengalaman baru (kosongkan jika tidak diubah): ")
	inputBytes, _ = reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].Pengalaman = input
	}

	fmt.Printf("Keahlian lama: %s\n", dataResume[indeks].Keahlian)
	fmt.Print("Keahlian baru (kosongkan jika tidak diubah): ")
	inputBytes, _ = reader.ReadString('\n')
	input = strings.TrimSpace(inputBytes)
	if input != "" {
		dataResume[indeks].Keahlian = input
	}

	fmt.Printf("Posisi dilamar lama: %s\n", dataResume[indeks].PosisiDilamar)
	fmt.Print("Posisi yang dilamar baru (kosongkan jika tidak diubah): ")
	var posisiBaru string
	posisiBaruBytes, _ := reader.ReadString('\n')
	posisiBaru = strings.TrimSpace(posisiBaruBytes)
	if posisiBaru != "" {
		dataResume[indeks].PosisiDilamar = posisiBaru
	}

	fmt.Print("Nama Perusahaan (untuk regenerate surat lamaran, kosongkan jika tidak diubah): ")
	var namaPerusahaan string
	namaPerusahaanBytes, _ := reader.ReadString('\n')
	namaPerusahaan = strings.TrimSpace(namaPerusahaanBytes)

	if namaPerusahaan != "" || posisiBaru != "" || input != "" {
		dataResume[indeks].SuratLamaran = generateSuratTemplate(namaPerusahaan, dataResume[indeks].Nama, dataResume[indeks].PosisiDilamar, dataResume[indeks].Keahlian)
		fmt.Println("Surat lamaran berhasil diupdate berdasarkan perubahan.")
	}

	fmt.Println("Data resume berhasil diedit.")
}

// mmenghapus data dari fungsi "TampilData" memalui value "ID"
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

// urutan data dari terkecil ke terbesar
func urutkanResumeAscending() {
	var i, pass int
	var temp Resume

	N := len(dataResume)

	pass = 1
	for pass < N {
		i = pass
		temp = dataResume[pass]

		for i > 0 && strings.ToLower(temp.Nama) < strings.ToLower(dataResume[i-1].Nama) {
			dataResume[i] = dataResume[i-1]
			i = i - 1
		}
		dataResume[i] = temp
		pass = pass + 1
	}
}

func urutkanResumeDescending() {
	N := len(dataResume)
	pass := 1
	for pass < N {
		idx := pass - 1
		for i := pass; i < N; i++ {
			if strings.ToLower(dataResume[i].Nama) > strings.ToLower(dataResume[idx].Nama) {
				idx = i
			}
		}
		dataResume[pass-1], dataResume[idx] = dataResume[idx], dataResume[pass-1]
		pass++
	}
}

// fungsi sebagai pembersih tampilan terminal
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

func pressEnterToContinue() {
	fmt.Println("\nTekan ENTER untuk melanjutkan...")
	var dummy string
	fmt.Scanln(&dummy)
}
