package main

import "fmt"

const NMAX = 100

type rekening struct {
	noRekening string
	nama       string
	saldo      float64
}
type TabRekening [NMAX]rekening

func TabelMenu() {
	fmt.Println("MENU DASBOARD ATM KONTROL AKSES")
	fmt.Println("1 MENAMBAHKAN DATA REKENING BARU")
	fmt.Println("2.HAPUS DATA REKENING")
	fmt.Println("3. MENGUBAH DATA REKENING NASABAH")
	fmt.Println("4. SETOR TARIK TUNAI")
	fmt.Println("5 MENCARI REKENING NASABAH BY SEQUENTIAL SEARCH")
	fmt.Println("6 MENGURUTKAN DATA NASABAH")
	fmt.Println("MENAMPILKAN SELURUH DATA ADMIN")

}
func validasilogin() bool {
	var username string
	var password string
	var i int
	var percobaan int
    percobaan = 0
	for i = 0; i < 3; i++ {
		percobaan = percobaan + 1
		fmt.Println("MENU LOGIN AKSES DASBOARD BANK")
		fmt.Println("USERNAME : ")
		fmt.Scan(&username)
		fmt.Println("PASSWORD : ")
		if username == "endra" && password == "telkomendra" {
			status = true
		}
	}
	if status == true && percobaan < 3 {
		statuslogin = true
	} else {
		statuslogin = true
	}
	return statuslogin
}
func inputdata(A *TabRekening, n *int) {
	fmt.Println("NAMA NASABAH")
	fmt.Scan(&A.nama)
	fmt.Println("SALDO YANG INGIN DI TAMBAHKAN")
	fmt.Scan(&A.saldo)
	fmt.Println("MASUKAN NOMOR REKENING NASABAH")
	fmt.Scan(&A.noRekening)
	*n = *n + 1
}
func deletedata(A *TabRekening, n *int,rekeningFound int) {
    var i int 
    var found bool
    found = false
    for repeat == "yes" {
    for i = 0;i <*n;i++{
        if A[i].noRekening == rekeningFound {
            for j = i;j < *n-1;j++{
                A[j] = A[j+1]
            }
            *n = *n - 1
            found = true
        }
        if found {
            fmt.Println("Data berhasil di hapus")
        } else {
            fmt.Println("Data berhasil di temukan")
        }
    }
    fmt.Println("Apakah anda ingin menghapus data lagi ? (yes/no)")
}
func main() {

}
