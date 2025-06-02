package main

// Kode dibuat oleh Rizky Febriyanto dan Orlando Silas Davinci Kambu IF-48-07

/*
"math/rand" & "time" nanti dia buat random pick templatenya gituu
"os/exec" & "runtime" sebagai fungsi dari clearscreen
"strings" buat struktur sebuah string yang terikat dengan array (digunakan untuk search dan sort)
*/
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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

func main() {
	landingPage()
}

// landing page awal
func landingPage() {
	clearScreen()
	fmt.Println("---------------------------------")
	fmt.Println(" Aplikasi Pembuat Resume & Lamaran ")
	fmt.Println("---------------------------------")
	fmt.Println("1. Daftar Akun Baru")
	fmt.Println("2. Masuk ke Aplikasi")
	fmt.Println("3. Keluar")
	fmt.Println("---------------------------------")
	fmt.Print("Pilih opsi: ")

	var choice int
	_, err := fmt.Scanln(&choice) // Menggunakan Scanln
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
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")

		var choice string
		_, err := fmt.Scanln(&choice) // Menggunakan Scanln untuk membaca pilihan
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
			return
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

	fmt.Println("-----------------------")
	fmt.Println(" Daftar Akun Baru ")
	fmt.Println("-----------------------")
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
			landingPage() // Kembali ke halaman utama
			return
		}
	}

	users = append(users, User{Username: username, Password: password})
	fmt.Println("\nAkun berhasil didaftarkan! Silakan login.")
	pressEnterToContinue()
	landingPage()
}

// fungsi login pengguna
func login() {
	var inputUsn, inputPass string
	clearScreen()
	fmt.Println("-----------------------")
	fmt.Println(" Masuk Aplikasi ")
	fmt.Println("-----------------------")
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

// kumpulan template surat lamaran yang nanti akan diinput oleh fungsi "TambahResume"
func GenerateSuratTemplate(nama, posisi, keahlian string) string {
	var templateList = [9]string{
		`Yth. HRD %s,

Dengan hormat,

Saya yang bertanda tangan di bawah ini, %s, bermaksud melamar posisi %s di perusahaan Anda. Saya memiliki keahlian dalam %s yang saya yakin akan bermanfaat bagi perusahaan.

Saya berharap dapat diberikan kesempatan untuk menunjukkan kompetensi saya lebih lanjut.

Hormat saya,
%s`,

		`Kepada Yth. Tim Rekrutmen %s,

Perkenalkan, saya %s. Dengan ini saya menyatakan minat saya untuk melamar posisi %s. Saya memiliki kemampuan di bidang %s yang telah saya kembangkan selama beberapa waktu.

Saya sangat menantikan kesempatan wawancara untuk mendiskusikan kontribusi saya.

Salam hormat,
%s`,

		`Halo %s,

Nama saya %s dan saya ingin melamar sebagai %s di perusahaan Anda. Keahlian saya di bidang %s membuat saya percaya diri untuk menjalankan tanggung jawab tersebut dengan baik.

Terima kasih atas perhatian Anda.

Hormat saya,
%s`,

		`Kepada HRD %s yang saya hormati,

Saya, %s, tertarik untuk mengisi posisi %s yang sedang dibuka. Saya memiliki pengalaman dan keahlian dalam %s yang dapat mendukung kinerja tim Anda.

Saya siap untuk mengikuti tahap seleksi lebih lanjut sesuai prosedur perusahaan.

Hormat saya,
%s`,

		`Yth. Manajer Perekrutan %s,

Nama saya %s dan saya ingin mengajukan lamaran untuk posisi %s. Dengan keahlian saya dalam %s, saya yakin dapat memberikan kontribusi positif dan solusi yang kreatif bagi perusahaan Anda.

Saya berharap dapat berdiskusi lebih lanjut melalui wawancara.

Hormat saya,
%s`,

		`Kepada %s Recruitment Team,

Saya, %s, bermaksud melamar posisi %s. Saya memiliki latar belakang dan keterampilan dalam bidang %s yang sesuai dengan kebutuhan posisi tersebut.

Saya sangat antusias dengan peluang untuk bergabung dan berkembang bersama perusahaan Anda.

Terima kasih,
%s`,

		`Yth. HRD %s,

Salam sejahtera,

Perkenalkan, saya %s. Saya sangat tertarik untuk melamar sebagai %s. Keahlian saya dalam %s menjadikan saya kandidat yang siap berkontribusi secara optimal.

Besar harapan saya untuk dapat bergabung dengan perusahaan Anda.

Hormat saya,
%s`,

		`Kepada Tim %s,

Saya adalah %s dan saya ingin melamar posisi %s. Keahlian saya di bidang %s telah terbukti dalam beberapa proyek sebelumnya, dan saya yakin dapat membawa nilai tambah.

Saya bersedia mengikuti seleksi sesuai ketentuan perusahaan.

Terima kasih,
%s`,

		`Yth. Bagian Rekrutmen %s,

Saya, %s, tertarik dengan posisi %s yang ditawarkan oleh perusahaan Anda. Keahlian saya dalam %s sangat relevan dengan tanggung jawab pekerjaan tersebut.

Saya siap memberikan yang terbaik jika diberikan kesempatan.

Salam hormat,
%s`,
	}

	// Inisialisasi random
	rand.Seed(time.Now().UnixNano())
	indeks := rand.Intn(len(templateList))
	return fmt.Sprintf(templateList[indeks], posisi, nama, posisi, keahlian, nama)
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

// menampilkan data hasil inputan dari "TambahResume" yang nanti akan digenerate oleh "GenerateSuratTemplate"
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

// mengedit data "TampilData" yang dibuat melalui "GenerateSuratTemplate" melalui "ID"
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

// urutaan data dari terbesar ke terkecil
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
