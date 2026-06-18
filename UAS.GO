package main

import "fmt"

const NMAX = 100

type Tagihan struct {
	Nama       string
	Nominal    int
	JatuhTempo string
	Kategori   string
	Status     string
}

var data [NMAX]Tagihan
var n int

// Data Dummy
func isiDataDummy() {
	data[0] = Tagihan{"Listrik", 250000, "06-25", "RumahTangga", "BelumLunas"}
	data[1] = Tagihan{"Air", 120000, "06-22", "RumahTangga", "Lunas"}
	data[2] = Tagihan{"Internet", 350000, "06-20", "Komunikasi", "BelumLunas"}
	data[3] = Tagihan{"Netflix", 65000, "06-18", "Hiburan", "Lunas"}
	data[4] = Tagihan{"BPJS", 150000, "06-15", "Kesehatan", "BelumLunas"}
	data[5] = Tagihan{"Spotify", 55000, "06-10", "Hiburan", "Lunas"}

	n = 6
}

// Menambah data tagihan
func tambahTagihan() {

	if n >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n=== Tambah Tagihan ===")

	fmt.Print("Nama Tagihan : ")
	fmt.Scan(&data[n].Nama)

	fmt.Print("Nominal : ")
	fmt.Scan(&data[n].Nominal)

	fmt.Print("Jatuh Tempo (Bulan-Tanggal) : ")
	fmt.Scan(&data[n].JatuhTempo)

	fmt.Print("Kategori : ")
	fmt.Scan(&data[n].Kategori)

	fmt.Print("Status (Lunas/BelumLunas) : ")
	fmt.Scan(&data[n].Status)

	n++
	fmt.Println("Data berhasil ditambahkan")
}

// Menampilkan data tagihan
func tampilTagihan() {

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Println("\n=== DAFTAR TAGIHAN ===")

	for i := 0; i < n; i++ {
		fmt.Println("----------------------")
		fmt.Println("No          :", i+1)
		fmt.Println("Nama        :", data[i].Nama)
		fmt.Println("Nominal     :", data[i].Nominal)
		fmt.Println("Jatuh Tempo :", data[i].JatuhTempo)
		fmt.Println("Kategori    :", data[i].Kategori)
		fmt.Println("Status      :", data[i].Status)
	}
}

// Mengubah data tagihan
func ubahTagihan() {

	var no int

	tampilTagihan()

	fmt.Print("Masukkan nomor yang ingin diubah : ")
	fmt.Scan(&no)

	if no >= 1 && no <= n {

		idx := no - 1

		fmt.Print("Nama Baru : ")
		fmt.Scan(&data[idx].Nama)

		fmt.Print("Nominal Baru : ")
		fmt.Scan(&data[idx].Nominal)

		fmt.Print("Jatuh Tempo Baru : ")
		fmt.Scan(&data[idx].JatuhTempo)

		fmt.Print("Kategori Baru : ")
		fmt.Scan(&data[idx].Kategori)

		fmt.Print("Status Baru : ")
		fmt.Scan(&data[idx].Status)

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Nomor tidak ditemukan")
	}
}

// Menghapus data tagihan
func hapusTagihan() {

	var no int

	tampilTagihan()

	fmt.Print("Masukkan nomor yang ingin dihapus : ")
	fmt.Scan(&no)

	if no >= 1 && no <= n {

		idx := no - 1

		for i := idx; i < n-1; i++ {
			data[i] = data[i+1]
		}

		n--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Nomor tidak ditemukan")
	}
}

// Ubah status pembayaran
func ubahStatusPembayaran() {

	var no int
	var status string

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	tampilTagihan()

	fmt.Print("Pilih nomor tagihan : ")
	fmt.Scan(&no)

	if no < 1 || no > n {
		fmt.Println("Nomor tidak valid")
		return
	}

	fmt.Print("Status Baru (Lunas/BelumLunas) : ")
	fmt.Scan(&status)

	if status != "Lunas" && status != "BelumLunas" {
		fmt.Println("Status tidak valid")
		return
	}

	data[no-1].Status = status

	fmt.Println("Status berhasil diperbarui")
}

// Ringkasan status pembayaran
func ringkasanStatus() {

	var lunas, belum int

	for i := 0; i < n; i++ {

		if data[i].Status == "Lunas" {
			lunas++
		} else {
			belum++
		}
	}

	fmt.Println("\n=== Ringkasan Pembayaran ===")
	fmt.Println("Lunas       :", lunas)
	fmt.Println("Belum Lunas :", belum)
}

// Sequential Search kategori
func sequentialSearchKategori() {

	var cari string
	var ketemu bool

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Print("Masukkan kategori yang dicari : ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {

		if data[i].Kategori == cari {

			fmt.Println("\nData Ditemukan")
			fmt.Println("Nama        :", data[i].Nama)
			fmt.Println("Nominal     :", data[i].Nominal)
			fmt.Println("Jatuh Tempo :", data[i].JatuhTempo)
			fmt.Println("Kategori    :", data[i].Kategori)
			fmt.Println("Status      :", data[i].Status)

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

// Selection Sort nama A-Z
func selectionSortNama() {

	var min int
	var temp Tagihan

	for i := 0; i < n-1; i++ {

		min = i

		for j := i + 1; j < n; j++ {

			if data[j].Nama < data[min].Nama {
				min = j
			}
		}

		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

// Binary Search nama tagihan
func binarySearchNama() {

	var cari string

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	selectionSortNama()

	fmt.Print("Masukkan nama tagihan : ")
	fmt.Scan(&cari)

	kiri := 0
	kanan := n - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if data[tengah].Nama == cari {

			fmt.Println("\nData ditemukan")
			fmt.Println("Nama        :", data[tengah].Nama)
			fmt.Println("Nominal     :", data[tengah].Nominal)
			fmt.Println("Jatuh Tempo :", data[tengah].JatuhTempo)
			fmt.Println("Kategori    :", data[tengah].Kategori)
			fmt.Println("Status      :", data[tengah].Status)

			return

		} else if cari < data[tengah].Nama {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

// Selection Sort jatuh tempo
func selectionSortTanggal() {

	var min int
	var temp Tagihan

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	for i := 0; i < n-1; i++ {

		min = i

		for j := i + 1; j < n; j++ {

			if data[j].JatuhTempo < data[min].JatuhTempo {
				min = j
			}
		}

		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}

	fmt.Println("Data berhasil diurutkan dengan Selection Sort")
	tampilTagihan()
}

// Insertion Sort jatuh tempo
func insertionSortTanggal() {

	var temp Tagihan
	var j int

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	for i := 1; i < n; i++ {

		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].JatuhTempo > temp.JatuhTempo {

			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan dengan Insertion Sort")
	tampilTagihan()
}

// Statistik tagihan
func statistik() {
	var totalBiaya int
	var jumlahLunas int
	var persentase float64

	if n == 0 {
		fmt.Println("Belum ada data")
		return
	}

	indeksTerdekat := -1

	for i := 0; i < n; i++ {

		totalBiaya += data[i].Nominal

		if data[i].Status == "Lunas" {
			jumlahLunas++
		} else {

			if indeksTerdekat == -1 || data[i].JatuhTempo < data[indeksTerdekat].JatuhTempo {
				indeksTerdekat = i
			}
		}
	}

	persentase = float64(jumlahLunas) / float64(n) * 100

	fmt.Println("\n=== STATISTIK TAGIHAN ===")
	fmt.Println("Total Biaya      :", totalBiaya)
	fmt.Printf("Persentase Lunas : %.2f%%\n", persentase)

	fmt.Println("---------------------------------")

	if indeksTerdekat != -1 {
		fmt.Println("Tagihan Terdekat yang Harus Dibayar:")
		fmt.Println("Nama        :", data[indeksTerdekat].Nama)
		fmt.Println("Jatuh Tempo :", data[indeksTerdekat].JatuhTempo)
		fmt.Println("Nominal     :", data[indeksTerdekat].Nominal)
	} else {
		fmt.Println("Semua tagihan sudah Lunas! Aman!")
	}
}

func main() {

	var pilihan int

	isiDataDummy()

	for {

		fmt.Println("\n+++ SIMTAB +++")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Tampilkan Tagihan")
		fmt.Println("3. Ubah Tagihan")
		fmt.Println("4. Hapus Tagihan")
		fmt.Println("5. Ubah Status Pembayaran")
		fmt.Println("6. Ringkasan Pembayaran")
		fmt.Println("7. Cari Kategori (Sequential Search)")
		fmt.Println("8. Cari Nama (Binary Search)")
		fmt.Println("9. Selection Sort Jatuh Tempo")
		fmt.Println("10. Insertion Sort Jatuh Tempo")
		fmt.Println("11. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahTagihan()

		case 2:
			tampilTagihan()

		case 3:
			ubahTagihan()

		case 4:
			hapusTagihan()

		case 5:
			ubahStatusPembayaran()

		case 6:
			ringkasanStatus()

		case 7:
			sequentialSearchKategori()

		case 8:
			binarySearchNama()

		case 9:
			selectionSortTanggal()

		case 10:
			insertionSortTanggal()

		case 11:
			statistik()

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}
