package main

import (
	"fmt"
)

const NMAX int = 100

type rekening struct {
	noRekening int
	nama       string
	saldo      float64
}

type TabRekening [NMAX]rekening

func tabelMenu() {
	fmt.Println("\n====================================")
	fmt.Println("      DASHBOARD ADMIN BANK")
	fmt.Println("====================================")
	fmt.Println("1. Tambah Rekening")
	fmt.Println("2. Hapus Rekening")
	fmt.Println("3. Ubah Data Rekening")
	fmt.Println("4. Setor Tunai")
	fmt.Println("5. Tarik Tunai")
	fmt.Println("6. Urutkan Berdasarkan Saldo")
	fmt.Println("7. Urutkan Berdasarkan Nama")
	fmt.Println("8. Cari Rekening")
	fmt.Println("9. Tampilkan Semua Data")
	fmt.Println("10. Keluar")
	fmt.Print("Pilih Menu : ")
}
func binarysearchNama(A *TabRekening, n int, nama string) {
	var mid, left, right int
	var found int
	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if nama < A[mid].nama {
			right = mid - 1
		} else if nama > A[mid].nama {
			left = mid + 1
		} else {
			found = mid
		}

	}
}
func validasiLogin(statusLogin *bool) {
	var username, password string
	var percobaan int

	*statusLogin = false
	percobaan = 0

	for percobaan < 3 && !(*statusLogin) {

		fmt.Println("\n===== LOGIN ADMIN =====")
		fmt.Print("Username : ")
		fmt.Scan(&username)

		fmt.Print("Password : ")
		fmt.Scan(&password)

		if username == "endra" && password == "telkomendra" {
			*statusLogin = true
		} else {
			fmt.Println("Username atau Password Salah")
			percobaan = percobaan + 1
		}
	}

	if *statusLogin {
		fmt.Println("Login Berhasil")
	} else {
		fmt.Println("Kesempatan Login Habis")
	}
}

func inputData(A *TabRekening, n *int) {

	if *n < NMAX {

		fmt.Print("Masukkan Nama Nasabah : ")
		fmt.Scan(&A[*n].nama)

		fmt.Print("Masukkan Nomor Rekening : ")
		fmt.Scan(&A[*n].noRekening)

		fmt.Print("Masukkan Saldo Awal : ")
		fmt.Scan(&A[*n].saldo)

		*n = *n + 1

		fmt.Println("Data Berhasil Ditambahkan")

	} else {
		fmt.Println("Kapasitas Penyimpanan Penuh")
	}
}

func sequentialSearch(A TabRekening, n int, noRekening int) {
	var ketemu int
	var k int

	ketemu = -1
	k = 0

	for ketemu == -1 && k < n {
		if A[k].noRekening == noRekening {
			ketemu = k
		}
		k = k + 1
	}
}

func hapusData(A *TabRekening, n *int) {
	var noRekening int
	var idx int
	var found bool
	var i int

	fmt.Print("Masukkan Nomor Rekening Yang Akan Dihapus : ")
	fmt.Scan(&noRekening)

	sequentialSearch(*A, *n, noRekening)

	if found {

		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n = *n - 1

		fmt.Println("Data Berhasil Dihapus")

	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func ubahData(A *TabRekening, n int) {
	var noRekening int
	var idx int
	var found bool

	fmt.Print("Masukkan Nomor Rekening Yang Akan Diubah : ")
	fmt.Scan(&noRekening)

	sequentialSearch(*A, n, noRekening)

	if found {

		fmt.Print("Nama Baru : ")
		fmt.Scan(&A[idx].nama)

		fmt.Print("Nomor Rekening Baru : ")
		fmt.Scan(&A[idx].noRekening)

		fmt.Print("Saldo Baru : ")
		fmt.Scan(&A[idx].saldo)

		fmt.Println("Data Berhasil Diubah")

	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func setorTunai(A *TabRekening, n int) {
	var noRekening int
	var nominal float64
	var idx int
	var found bool

	fmt.Print("Masukkan Nomor Rekening : ")
	fmt.Scan(&noRekening)

	sequentialSearch(*A, n, noRekening)

	if found {

		fmt.Print("Masukkan Nominal Setor : ")
		fmt.Scan(&nominal)

		A[idx].saldo = A[idx].saldo + nominal

		fmt.Println("Setor Tunai Berhasil")
		fmt.Println("Saldo Sekarang :", A[idx].saldo)

	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func tarikTunai(A *TabRekening, n int) {
	var noRekening int
	var nominal float64
	var idx int
	var found bool

	fmt.Print("Masukkan Nomor Rekening : ")
	fmt.Scan(&noRekening)

	sequentialSearch(*A, n, noRekening)

	if found {

		fmt.Print("Masukkan Nominal Tarik : ")
		fmt.Scan(&nominal)

		if nominal <= A[idx].saldo {

			A[idx].saldo = A[idx].saldo - nominal

			fmt.Println("Tarik Tunai Berhasil")
			fmt.Println("Saldo Sekarang :", A[idx].saldo)

		} else {
			fmt.Println("Saldo Tidak Mencukupi")
		}

	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func insertionSortBySaldo(A *TabRekening, n int) {
	var pass, i int
	var temp rekening

	pass = 1

	for pass < n-1 {

		temp = A[pass]
		i = pass

		for i > 0 && temp.saldo > A[i-1].saldo {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
		pass = pass + 1
	}

	fmt.Println("Data Berhasil Diurutkan Berdasarkan Saldo")
}
func insertionSortByNama(A *TabRekening, n int) {
	var i, pass int
	var temp rekening

	pass = 1

	for pass <= n-1 {
		temp = A[pass]
		i = pass
		for pass > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
	fmt.Println("Data Berhasil Diurutkan Berdasarkan Nama")
}
func cariRekening(A TabRekening, n int) {
	var noRekening int
	var idx int
	var found bool

	fmt.Print("Masukkan Nomor Rekening : ")
	fmt.Scan(&noRekening)

	sequentialSearch(A, n, noRekening)

	if found {

		fmt.Println("\n===== DATA DITEMUKAN =====")
		fmt.Println("Nama         :", A[idx].nama)
		fmt.Println("No Rekening  :", A[idx].noRekening)
		fmt.Println("Saldo        :", A[idx].saldo)

	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func tampilData(A TabRekening, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum Ada Data")
	} else {

		fmt.Println("\n========== DATA NASABAH ==========")

		for i = 0; i < n; i++ {

			fmt.Println("--------------------------------")
			fmt.Println("Data Ke-", i+1)
			fmt.Println("Nama         :", A[i].nama)
			fmt.Println("No Rekening  :", A[i].noRekening)
			fmt.Println("Saldo        :", A[i].saldo)
		}
	}
}

func main() {
	var nama string
	var data TabRekening
	var n int
	var pilihan int
	var statusLogin bool
	var selesai bool

	selesai = false

	validasiLogin(&statusLogin)

	if statusLogin {

		for !selesai {

			tabelMenu()
			fmt.Scan(&pilihan)

			switch pilihan {

			case 1:
				inputData(&data, &n)

			case 2:
				hapusData(&data, &n)

			case 3:
				ubahData(&data, n)

			case 4:
				setorTunai(&data, n)

			case 5:
				tarikTunai(&data, n)

			case 6:
				insertionSortBySaldo(&data, n)

			case 7:
				insertionSortByNama(&data, n)
				binarysearchNama(&data, n, nama)

			case 8:
				tampilData(data, n)

			case 9:
				tampilData(data, n)

			case 10:
				selesai = true
				fmt.Println("Terima Kasih")
				fmt.Println("BLOG PRIBADI https://endrasg.web.id/ ")

			default:
				fmt.Println("Menu Tidak Tersedia")
			}
		}

	} else {
		fmt.Println("Akses Ditolak")
	}
}
